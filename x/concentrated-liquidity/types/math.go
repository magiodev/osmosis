package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/osmomath"
)

// tickToSqrtPrice takes the tick index and returns the corresponding sqrt of the price
func TickToSqrtPrice(tickIndex sdk.Int) (sdk.Dec, error) {
	price, err := sdk.NewDecWithPrec(10001, 4).Power(tickIndex.Uint64()).ApproxSqrt()
	if err != nil {
		return sdk.Dec{}, err
	}

	return price, nil
}

// TODO: implement this
//
//	func (k Keeper) getTickAtSqrtRatio(sqrtPrice sdk.Dec) sdk.Int {
//		return sdk.Int{}
//	}
//
// PriceToTick takes a price and returns the corresponding tick index
func PriceToTick(price sdk.Dec) sdk.Int {
	logOfPrice := osmomath.BigDecFromSDKDec(price).LogBase2()
	logInt := osmomath.NewDecWithPrec(10001, 4)
	tick := logOfPrice.Quo(logInt.LogBase2())
	return tick.SDKDec().TruncateInt()
}

// liquidity0 takes an amount of asset0 in the pool as well as the sqrtpCur and the nextPrice
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// liquidity0 = amount0 * (sqrtPriceA * sqrtPriceB) / (sqrtPriceB - sqrtPriceA)
func Liquidity0(amount sdk.Int, sqrtPriceA, sqrtPriceB sdk.Dec) sdk.Dec {
	if sqrtPriceA.GT(sqrtPriceB) {
		sqrtPriceA, sqrtPriceB = sqrtPriceB, sqrtPriceA
	}
	product := sqrtPriceA.Mul(sqrtPriceB)
	diff := sqrtPriceB.Sub(sqrtPriceA)
	return amount.ToDec().Mul(product).Quo(diff)
}

// liquidity1 takes an amount of asset1 in the pool as well as the sqrtpCur and the nextPrice
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// liquidity1 = amount1 / (sqrtPriceB - sqrtPriceA)
func Liquidity1(amount sdk.Int, sqrtPriceA, sqrtPriceB sdk.Dec) sdk.Dec {
	if sqrtPriceA.GT(sqrtPriceB) {
		sqrtPriceA, sqrtPriceB = sqrtPriceB, sqrtPriceA
	}
	diff := sqrtPriceB.Sub(sqrtPriceA)
	return amount.ToDec().Quo(diff)
}

// calcAmount0 takes the asset with the smaller liquidity in the pool as well as the sqrtpCur and the nextPrice and calculates the amount of asset 0
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// calcAmount0Delta = (liquidity * (sqrtPriceB - sqrtPriceA)) / (sqrtPriceB * sqrtPriceA)
func CalcAmount0Delta(liq, sqrtPriceA, sqrtPriceB sdk.Dec, roundUp bool) sdk.Dec {
	if sqrtPriceA.GT(sqrtPriceB) {
		sqrtPriceA, sqrtPriceB = sqrtPriceB, sqrtPriceA
	}
	diff := sqrtPriceB.Sub(sqrtPriceA)
	denom := sqrtPriceA.Mul(sqrtPriceB)
	// if calculating for amountIn, we round up
	// if calculating for amountOut, we don't round at all
	// this is to prevent removing more from the pool than expected due to rounding
	// example: we calculate 1000000.9999999 uusdc (~$1) amountIn and 2000000.999999 uosmo amountOut
	// we would want the user to put in 1000001 uusdc rather than 1000000 uusdc to ensure we are charging enough for the amount they are removing
	// additionally, without rounding, there exists cases where the swapState.amountSpecifiedRemaining.GT(sdk.ZeroDec()) for loop within
	// the CalcOut/In functions never actually reach zero due to dust that would have never gotten counted towards the amount (numbers after the 10^6 place)
	if roundUp {
		return liq.Mul(diff.Quo(denom)).Ceil()
	}
	return liq.Mul(diff.Quo(denom))
}

// calcAmount1 takes the asset with the smaller liquidity in the pool as well as the sqrtpCur and the nextPrice and calculates the amount of asset 1
// sqrtPriceA is the smaller of sqrtpCur and the nextPrice
// sqrtPriceB is the larger of sqrtpCur and the nextPrice
// calcAmount1Delta = liq * (sqrtPriceB - sqrtPriceA)
func CalcAmount1Delta(liq, sqrtPriceA, sqrtPriceB sdk.Dec, roundUp bool) sdk.Dec {
	if sqrtPriceA.GT(sqrtPriceB) {
		sqrtPriceA, sqrtPriceB = sqrtPriceB, sqrtPriceA
	}
	diff := sqrtPriceB.Sub(sqrtPriceA)
	// if calculating for amountIn, we round up
	// if calculating for amountOut, we don't round at all
	// this is to prevent removing more from the pool than expected due to rounding
	// example: we calculate 1000000.9999999 uusdc (~$1) amountIn and 2000000.999999 uosmo amountOut
	// we would want the used to put in 1000001 uusdc rather than 1000000 uusdc to ensure we are charging enough for the amount they are removing
	// additionally, without rounding, there exists cases where the swapState.amountSpecifiedRemaining.GT(sdk.ZeroDec()) for loop within
	// the CalcOut/In functions never actually reach zero due to dust that would have never gotten counted towards the amount (numbers after the 10^6 place)
	if roundUp {
		return liq.Mul(diff).Ceil()
	}
	return liq.Mul(diff)
}

// computeSwapStep calculates the amountIn, amountOut, and the next sqrtPrice given current price, price target, tick liquidity, and amount available to swap
// lte is reference to "less than or equal", which determines if we are moving left or right of the current price to find the next initialized tick with liquidity
func ComputeSwapStep(sqrtPriceCurrent, sqrtPriceTarget, liquidity, amountRemaining sdk.Dec, zeroForOne bool) (sqrtPriceNext, amountIn, amountOut sdk.Dec) {
	// zeroForOne = sqrtPriceCurrent.GTE(sqrtPriceTarget)
	if zeroForOne {
		amountIn = CalcAmount0Delta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, false)
	} else {
		amountIn = CalcAmount1Delta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, false)
	}
	if amountRemaining.GTE(amountIn) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		sqrtPriceNext = GetNextSqrtPriceFromInput(sqrtPriceCurrent, liquidity, amountRemaining, zeroForOne)
	}

	if zeroForOne {
		amountIn = CalcAmount0Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
		amountOut = CalcAmount1Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
	} else {
		amountIn = CalcAmount1Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
		amountOut = CalcAmount0Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
	}

	return sqrtPriceNext, amountIn, amountOut
}

func GetNextSqrtPriceFromInput(sqrtPriceCurrent, liquidity, amountRemaining sdk.Dec, zeroForOne bool) (sqrtPriceNext sdk.Dec) {
	if zeroForOne {
		sqrtPriceNext = GetNextSqrtPriceFromAmount0RoundingUp(sqrtPriceCurrent, liquidity, amountRemaining)
	} else {
		sqrtPriceNext = GetNextSqrtPriceFromAmount1RoundingDown(sqrtPriceCurrent, liquidity, amountRemaining)
	}
	return sqrtPriceNext
}

// getNextSqrtPriceFromAmount0RoundingUp utilizes the current squareRootPrice, liquidity of denom0, and amount of denom0 that still needs
// to be swapped in order to determine the next squareRootPrice
// if (amountRemaining * sqrtPriceCurrent) / amountRemaining  == sqrtPriceCurrent AND (liquidity) + (amountRemaining * sqrtPriceCurrent) >= (liquidity)
// sqrtPriceNext = (liquidity * sqrtPriceCurrent) / ((liquidity) + (amountRemaining * sqrtPriceCurrent))
// else
// sqrtPriceNext = ((liquidity)) / (((liquidity) / (sqrtPriceCurrent)) + (amountRemaining))
func GetNextSqrtPriceFromAmount0RoundingUp(sqrtPriceCurrent, liquidity, amountRemaining sdk.Dec) (sqrtPriceNext sdk.Dec) {
	numerator := liquidity
	product := amountRemaining.Mul(sqrtPriceCurrent)

	if product.Quo(amountRemaining).Equal(sqrtPriceCurrent) {
		denominator := numerator.Add(product)
		if denominator.GTE(numerator) {
			numerator = numerator.Mul(sqrtPriceCurrent)
			sqrtPriceNext = numerator.QuoRoundUp(denominator)
			return sqrtPriceNext
		}
	}
	denominator := numerator.Quo(sqrtPriceCurrent).Add(amountRemaining)
	sqrtPriceNext = numerator.QuoRoundUp(denominator)
	return sqrtPriceNext
}

// getNextSqrtPriceFromAmount1RoundingDown utilizes the current squareRootPrice, liquidity of denom1, and amount of denom1 that still needs
// to be swapped in order to determine the next squareRootPrice
// sqrtPriceNext = sqrtPriceCurrent + (amount1Remaining / liquidity1)
func GetNextSqrtPriceFromAmount1RoundingDown(sqrtPriceCurrent, liquidity, amountRemaining sdk.Dec) (sqrtPriceNext sdk.Dec) {
	return sqrtPriceCurrent.Add(amountRemaining.Quo(liquidity))
}

// getLiquidityFromAmounts takes the current sqrtPrice and the sqrtPrice for the upper and lower ticks as well as the amounts of asset0 and asset1
// in return, liquidity is calculated from these inputs
func GetLiquidityFromAmounts(sqrtPrice, sqrtPriceA, sqrtPriceB sdk.Dec, amount0, amount1 sdk.Int) (liquidity sdk.Dec) {
	if sqrtPriceA.GT(sqrtPriceB) {
		sqrtPriceA, sqrtPriceB = sqrtPriceB, sqrtPriceA
	}
	if sqrtPrice.LTE(sqrtPriceA) {
		liquidity = Liquidity0(amount0, sqrtPriceA, sqrtPriceB)
	} else if sqrtPrice.LTE(sqrtPriceB) {
		liquidity0 := Liquidity0(amount0, sqrtPrice, sqrtPriceB)
		liquidity1 := Liquidity1(amount1, sqrtPrice, sqrtPriceA)
		liquidity = sdk.MinDec(liquidity0, liquidity1)
	} else {
		liquidity = Liquidity1(amount1, sqrtPriceB, sqrtPriceA)
	}
	return liquidity
}

func AddLiquidity(liquidityA, liquidityB sdk.Dec) (finalLiquidity sdk.Dec) {
	if liquidityB.LT(sdk.ZeroDec()) {
		return liquidityA.Sub(liquidityB.Abs())
	}
	return liquidityA.Add(liquidityB)
}
