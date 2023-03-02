package swapstrategy

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/internal/math"
	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
)

// oneForZeroStrategy implements the swapStrategy interface.
// This implementation assumes that we are swapping token 1 for
// token 0 and performs calculations accordingly.
//
// With this strategy, we are moving to the right of the current
// tick index and square root price.
type oneForZeroStrategy struct {
	sqrtPriceLimit sdk.Dec
	storeKey       sdk.StoreKey
	swapFee        sdk.Dec
}

var _ swapStrategy = (*oneForZeroStrategy)(nil)

// GetSqrtTargetPrice returns the target square root price given the next tick square root price.
// If the given nextTickSqrtPrice is greater than the sqrt price limit, the sqrt price limit is returned.
// Otherwise, the input nextTickSqrtPrice is returned.
func (s oneForZeroStrategy) GetSqrtTargetPrice(nextTickSqrtPrice sdk.Dec) sdk.Dec {
	if nextTickSqrtPrice.GT(s.sqrtPriceLimit) {
		return s.sqrtPriceLimit
	}
	return nextTickSqrtPrice
}

func (s oneForZeroStrategy) ComputeSwapStepOutGivenIn(sqrtPriceCurrent, sqrtPriceTarget, liquidity, amountOneInRemaining sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec) {
	// Estimate the amount of token one needed until the target sqrt price is reached.
	amountOneIn := math.CalcAmount1Delta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, true) // N.B.: if this is false, causes infinite loop

	// Calculate sqrtPriceNext on the amount of token remaining after fee.
	amountOneInRemainingLessFee := amountOneInRemaining.Mul(sdk.OneDec().Sub(s.swapFee))

	var sqrtPriceNext sdk.Dec
	// If have more of the amount remaining after fee than estimated until target,
	// bound the next sqrtPriceNext by the target sqrt price.
	if amountOneInRemainingLessFee.GTE(amountOneIn) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		// Otherwise, compute the next sqrt price based on the amount remaining after fee.
		sqrtPriceNext = math.GetNextSqrtPriceFromAmount1InRoundingDown(sqrtPriceCurrent, liquidity, amountOneInRemainingLessFee)
	}

	hasReachedTarget := sqrtPriceTarget == sqrtPriceNext

	// If the sqrt price target was not reached, recalculate how much of the amount remaining after fee was needed
	// to complete the swap step. This implies that some of the amount remaining after fee is left over after the
	// current swap step.
	if !hasReachedTarget {
		amountOneIn = math.CalcAmount1Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true) // N.B.: if this is false, causes infinite loop
	}

	// Calculate the amount of the other token given the sqrt price range.
	amountZeroOut := math.CalcAmount0Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)

	// Handle fees.
	// Note that fee is always charged on the amount in.
	feeChargeTotal := computeFeeChargePerSwapStepOutGivenIn(sqrtPriceCurrent, hasReachedTarget, amountOneIn, amountOneInRemaining, s.swapFee)

	fmt.Println("sqrtPriceCurrent", sqrtPriceCurrent)
	fmt.Println("sqrtPriceNext", sqrtPriceNext)
	fmt.Println("amountOneIn", amountOneIn)
	fmt.Println("amountZeroOut", amountZeroOut)

	return sqrtPriceNext, amountOneIn, amountZeroOut, feeChargeTotal
}

// ComputeSwapStep calculates the next sqrt price, the new amount remaining, the amount of the token other than remaining given current price, and total fee charge.
//
// zeroForOneStrategy assumes moving to the left of the current square root price.
// amountRemaining is the amount of token in when swapping out given in and token out when swapping in given out.
// amountRemaining is token zero.
// amountZero is token out when swapping in given out and token in when swapping out given in.
// amountOne is token in when swapping in given out and token out when swapping out given in.
// TODO: improve tests
func (s oneForZeroStrategy) ComputeSwapStepInGivenOut(sqrtPriceCurrent, sqrtPriceTarget, liquidity, amountRemainingOut sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec) {
	// Estimate the amount of token zero needed until the target sqrt price is reached.
	// N.B.: contrary to out given in, we do not round up because we do not want to exceed the initial amount out at the end.
	amountZeroOut := math.CalcAmount0Delta(liquidity, sqrtPriceTarget, sqrtPriceCurrent, false)

	// Calculate sqrtPriceNext on the amount of token remaining. Note that the
	// fee is not charged as amountRemaining is amountOut, and we only charge fee on
	// amount in.
	var sqrtPriceNext sdk.Dec
	// If have more of the amount remaining after fee than estimated until target,
	// bound the next sqrtPriceNext by the target sqrt price.
	if amountRemainingOut.GTE(amountZeroOut) {
		sqrtPriceNext = sqrtPriceTarget
	} else {
		// Otherwise, compute the next sqrt price based on the amount remaining after fee.
		sqrtPriceNext = math.GetNextSqrtPriceFromAmount0OutRoundingUp(sqrtPriceCurrent, liquidity, amountRemainingOut)
	}

	hasReachedTarget := sqrtPriceTarget == sqrtPriceNext

	// If the sqrt price target was not reached, recalculate how much of the amount remaining after fee was needed
	// to complete the swap step. This implies that some of the amount remaining after fee is left over after the
	// current swap step.
	if !hasReachedTarget {
		// N.B.: contrary to out given in, we do not round up because we do not want to exceed the initial amount out at the end.
		amountZeroOut = math.CalcAmount0Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, false)
	}

	// Calculate the amount of the other token given the sqrt price range.
	amountOneIn := math.CalcAmount1Delta(liquidity, sqrtPriceNext, sqrtPriceCurrent, true)

	// Handle fees.
	// Note that fee is always charged on the amount in.
	feeChargeTotal := amountOneIn.Mul(s.swapFee).Quo(sdk.OneDec().Sub(s.swapFee))

	return sqrtPriceNext, amountZeroOut, amountOneIn, feeChargeTotal
}

// InitializeTickValue returns the initial tick value for computing swaps based
// on the actual current tick.
//
// oneForZeroStrategy assumes moving to the right of the current square root price.
// As a result, we use forward iterator in NextInitializedTick to find the next
// tick to the left of current. The end cursor for forward iteration is inclusive.
// Therefore, this method is, essentially a no-op. The logic is reversed for
// zeroForOneStrategy where we use reverse iterator and have to add one to
// the input. Therefore, we define this method to account for different strategies.
func (s oneForZeroStrategy) InitializeTickValue(currentTick sdk.Int) sdk.Int {
	return currentTick
}

// NextInitializedTick returns the next initialized tick index based on the
// provided tickindex. If no initialized tick exists, <0, false>
// will be returned.
//
// oneForZerostrategy searches for the next tick to the right of the current tickIndex.
func (s oneForZeroStrategy) NextInitializedTick(ctx sdk.Context, poolId uint64, tickIndex int64) (next sdk.Int, initialized bool) {
	store := ctx.KVStore(s.storeKey)

	// Construct a prefix store with a prefix of <TickPrefix | poolID>, allowing
	// us to retrieve the next initialized tick without having to scan all ticks.
	prefixBz := types.KeyTickPrefix(poolId)
	prefixStore := prefix.NewStore(store, prefixBz)

	startKey := types.TickIndexToBytes(tickIndex)

	iter := prefixStore.Iterator(startKey, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		// Since, we constructed our prefix store with <TickPrefix | poolID>, the
		// key is the encoding of a tick index.
		tick, err := types.TickIndexFromBytes(iter.Key())
		if err != nil {
			panic(fmt.Errorf("invalid tick index (%s): %v", string(iter.Key()), err))
		}

		if tick > tickIndex {
			return sdk.NewInt(tick), true
		}
	}
	return sdk.ZeroInt(), false
}

// SetLiquidityDeltaSign sets the liquidity delta sign for the given liquidity delta.
// This is called when consuming all liquidity.
// When a position is created, we add liquidity to lower tick
// and subtract from the upper tick to reflect that this new
// liquidity would be added when the price crosses the lower tick
// going up, and subtracted when the price crosses the upper tick
// going up. As a result, the sign depend on the direction we are moving.
//
// oneForZeroStrategy assumes moving to the right of the current square root price.
// When we move to the right, we must be crossing lower ticks first where
// liqudiity delta tracks the amount of liquidity being added. So the sign must be
// positive.
func (s oneForZeroStrategy) SetLiquidityDeltaSign(deltaLiquidity sdk.Dec) sdk.Dec {
	return deltaLiquidity
}

// ValidatePriceLimit validates the given square root price limit
// given the square root price.
//
// oneForZeroStrategy assumes moving to the right of the current square root price.
// Therefore, the following invariant must hold:
// current square root price <= sqrtPriceLimit <= types.MaxSqrtRatio
func (s oneForZeroStrategy) ValidatePriceLimit(sqrtPriceLimit, currentSqrtPrice sdk.Dec) error {
	// check that the price limit is above the current sqrt price but lower than the maximum sqrt ratio since we are swapping asset1 for asset0
	if sqrtPriceLimit.LT(currentSqrtPrice) || sqrtPriceLimit.GT(types.MaxSqrtRatio) {
		return types.InvalidPriceLimitError{SqrtPriceLimit: sqrtPriceLimit, LowerBound: currentSqrtPrice, UpperBound: types.MaxSqrtRatio}
	}
	return nil
}
