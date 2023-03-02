package swapstrategy_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/internal/swapstrategy"
	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
)

func (suite *StrategyTestSuite) TestGetSqrtTargetPrice_OneForZero() {
	tests := map[string]struct {
		sqrtPriceLimit    sdk.Dec
		nextTickSqrtPrice sdk.Dec
		expectedResult    sdk.Dec
	}{
		"nextTickSqrtPrice == sqrtPriceLimit -> returns either": {
			sqrtPriceLimit:    sdk.OneDec(),
			nextTickSqrtPrice: sdk.OneDec(),
			expectedResult:    sdk.OneDec(),
		},
		"nextTickSqrtPrice > sqrtPriceLimit -> sqrtPriceLimit": {
			sqrtPriceLimit:    three,
			nextTickSqrtPrice: four,
			expectedResult:    three,
		},
		"nextTickSqrtPrice < sqrtPriceLimit -> nextTickSqrtPrice": {
			sqrtPriceLimit:    five,
			nextTickSqrtPrice: two,
			expectedResult:    two,
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()

			sut := swapstrategy.New(false, tc.sqrtPriceLimit, suite.App.GetKey(types.ModuleName), sdk.ZeroDec())

			actualSqrtTargetPrice := sut.GetSqrtTargetPrice(tc.nextTickSqrtPrice)

			suite.Require().Equal(tc.expectedResult, actualSqrtTargetPrice)

		})
	}
}

func (suite *StrategyTestSuite) TestComputeSwapStepOutGivenIn_OneForZero() {
	tokenOutErrTolerance := osmomath.ErrTolerance{
		AdditiveTolerance: sdk.SmallestDec().MulInt64(200),
	}

	var (
		sqrtPriceCurrent = defaultSqrtPriceLower
		sqrtPriceNext    = defaultSqrtPriceUpper
		// Note, we take a ceiling here
		defaultAmountOne = defaultAmountOne.Ceil()

		// sqrt_price_current + token_in / liquidity
		sqrtPriceTargetNotReached = sdk.MustNewDecFromStr("74.161984805609412580")
		// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
		amountZeroTargetNotReached = sdk.MustNewDecFromStr("998976.600312992416750225")
	)

	tests := map[string]struct {
		sqrtPriceCurrent     sdk.Dec
		sqrtPriceTarget      sdk.Dec
		liquidity            sdk.Dec
		amountOneInRemaining sdk.Dec
		swapFee              sdk.Dec

		expectedSqrtPriceNext        sdk.Dec
		expectedNewAmountRemainingIn sdk.Dec
		expectedAmountOut            sdk.Dec
		expectedFeeChargeTotal       sdk.Dec

		expectError error
	}{
		"1: no fee - reach target": {
			sqrtPriceCurrent:     sqrtPriceCurrent,
			sqrtPriceTarget:      sqrtPriceNext,
			liquidity:            defaultLiquidity,
			amountOneInRemaining: defaultAmountOne,
			swapFee:              sdk.ZeroDec(),

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountOne,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
			expectedAmountOut:      defaultAmountZero,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"2: no fee - do not reach target": {
			sqrtPriceCurrent:     sqrtPriceCurrent,
			sqrtPriceTarget:      sqrtPriceNext,
			liquidity:            defaultLiquidity,
			amountOneInRemaining: defaultAmountOne.Sub(sdk.NewDec(100)),
			swapFee:              sdk.ZeroDec(),

			// sqrt_price_current + token_in / liquidity
			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountOne.Sub(sdk.NewDec(100)),
			expectedAmountOut:            amountZeroTargetNotReached,
			expectedFeeChargeTotal:       sdk.ZeroDec(),
		},
		"3: 3% fee - reach target": {
			sqrtPriceCurrent:     sqrtPriceCurrent,
			sqrtPriceTarget:      sqrtPriceNext,
			liquidity:            defaultLiquidity,
			amountOneInRemaining: defaultAmountOne.Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:              defaultFee,

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountOne,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
			expectedAmountOut:      defaultAmountZero,
			expectedFeeChargeTotal: defaultAmountOne.Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
		"4: 3% fee - do not reach target": {
			sqrtPriceCurrent:     sqrtPriceCurrent,
			sqrtPriceTarget:      sqrtPriceNext,
			liquidity:            defaultLiquidity,
			amountOneInRemaining: defaultAmountOne.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:              defaultFee,

			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountOne.Sub(sdk.NewDec(100)),
			expectedAmountOut:            amountZeroTargetNotReached,
			expectedFeeChargeTotal:       defaultAmountOne.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()
			strategy := swapstrategy.New(false, types.MaxSqrtRatio, suite.App.GetKey(types.ModuleName), tc.swapFee)

			sqrtPriceNext, newAmountRemainingOneIn, amountZeroOut, feeChargeTotal := strategy.ComputeSwapStepOutGivenIn(tc.sqrtPriceCurrent, tc.sqrtPriceTarget, tc.liquidity, tc.amountOneInRemaining)

			suite.Require().Equal(tc.expectedSqrtPriceNext, sqrtPriceNext)
			suite.Require().Equal(tc.expectedNewAmountRemainingIn, newAmountRemainingOneIn)
			suite.Require().Equal(0,
				tokenOutErrTolerance.CompareBigDec(
					osmomath.BigDecFromSDKDec(tc.expectedAmountOut),
					osmomath.BigDecFromSDKDec(amountZeroOut),
				),
				fmt.Sprintf("expected (%s), actual (%s)", tc.expectedAmountOut, amountZeroOut))
			suite.Require().Equal(tc.expectedFeeChargeTotal, feeChargeTotal)
		})
	}
}

func (suite *StrategyTestSuite) TestComputeSwapStepInGivenOut_OneForZero() {
	tokenOutErrTolerance := osmomath.ErrTolerance{
		AdditiveTolerance: sdk.SmallestDec().MulInt64(200),
	}

	smallestErrTolerance := osmomath.ErrTolerance{
		AdditiveTolerance: sdk.SmallestDec(),
	}

	var (
		sqrtPriceCurrent = defaultSqrtPriceLower
		sqrtPriceNext    = defaultSqrtPriceUpper
		// Note, we take a ceiling here.
		defaultAmountOne = defaultAmountOne.Ceil()

		// sqrtPriceCurrent + amount out / liquidity
		sqrtPriceTargetNotReached = sdk.MustNewDecFromStr("74.161984805609412580")
		// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
		amountOneTargetNotReached = sdk.MustNewDecFromStr("998976.600312992416750195")
	)

	tests := map[string]struct {
		sqrtPriceCurrent      sdk.Dec
		sqrtPriceTarget       sdk.Dec
		liquidity             sdk.Dec
		amountOneOutRemaining sdk.Dec
		swapFee               sdk.Dec

		expectedSqrtPriceNext         sdk.Dec
		expectedNewAmountRemainingOut sdk.Dec
		expectedAmountZeroIn          sdk.Dec
		expectedFeeChargeTotal        sdk.Dec

		expectError error
	}{
		"1: no fee - reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountOneOutRemaining: defaultAmountOne,
			swapFee:               sdk.ZeroDec(),

			expectedSqrtPriceNext:         sqrtPriceNext,
			expectedNewAmountRemainingOut: defaultAmountOne,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent)
			expectedAmountZeroIn:   defaultAmountZero,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"2: no fee - do not reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountOneOutRemaining: defaultAmountOne.Sub(sdk.NewDec(100)),
			swapFee:               sdk.ZeroDec(),

			expectedSqrtPriceNext:         sqrtPriceTargetNotReached,
			expectedNewAmountRemainingOut: defaultAmountOne.Sub(sdk.NewDec(100)),

			expectedAmountZeroIn:   amountOneTargetNotReached,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"3: 3% fee - reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountOneOutRemaining: defaultAmountOne.Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:               defaultFee,

			expectedSqrtPriceNext:         sqrtPriceNext,
			expectedNewAmountRemainingOut: defaultAmountOne,
			expectedAmountZeroIn:          defaultAmountZero,
			expectedFeeChargeTotal:        defaultAmountZero.Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
		"4: 3% fee - do not reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountOneOutRemaining: defaultAmountOne.Sub(sdk.NewDec(100)),
			swapFee:               defaultFee,

			expectedSqrtPriceNext:         sqrtPriceTargetNotReached,
			expectedNewAmountRemainingOut: defaultAmountOne.Sub(sdk.NewDec(100)),
			expectedAmountZeroIn:          amountOneTargetNotReached,
			expectedFeeChargeTotal:        amountOneTargetNotReached.Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()
			strategy := swapstrategy.New(true, types.MaxSqrtRatio, suite.App.GetKey(types.ModuleName), tc.swapFee)

			sqrtPriceNext, newAmountRemainingOneOut, amountZeroIn, feeChargeTotal := strategy.ComputeSwapStepInGivenOut(tc.sqrtPriceCurrent, tc.sqrtPriceTarget, tc.liquidity, tc.amountOneOutRemaining)

			suite.Require().Equal(tc.expectedSqrtPriceNext, sqrtPriceNext)
			suite.Require().Equal(tc.expectedNewAmountRemainingOut, newAmountRemainingOneOut)
			suite.Require().Equal(0,
				tokenOutErrTolerance.CompareBigDec(
					osmomath.BigDecFromSDKDec(tc.expectedAmountZeroIn),
					osmomath.BigDecFromSDKDec(amountZeroIn),
				),
				fmt.Sprintf("expected (%s), actual (%s)", tc.expectedAmountZeroIn, amountZeroIn))

			suite.Require().Equal(0,
				smallestErrTolerance.CompareBigDec(
					osmomath.BigDecFromSDKDec(tc.expectedFeeChargeTotal),
					osmomath.BigDecFromSDKDec(feeChargeTotal),
				),
				fmt.Sprintf("expected (%s), actual (%s)", tc.expectedAmountZeroIn, amountZeroIn))
		})
	}
}
