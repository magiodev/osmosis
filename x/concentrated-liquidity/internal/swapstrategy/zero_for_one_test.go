package swapstrategy_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/internal/swapstrategy"
	"github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
)

func (suite *StrategyTestSuite) TestGetSqrtTargetPrice_ZeroForOne() {
	var (
		two   = sdk.NewDec(2)
		three = sdk.NewDec(2)
		four  = sdk.NewDec(4)
		five  = sdk.NewDec(5)
	)

	tests := map[string]struct {
		isZeroForOne      bool
		sqrtPriceLimit    sdk.Dec
		nextTickSqrtPrice sdk.Dec
		expectedResult    sdk.Dec
	}{
		"nextTickSqrtPrice == sqrtPriceLimit -> returns either": {
			sqrtPriceLimit:    sdk.OneDec(),
			nextTickSqrtPrice: sdk.OneDec(),
			expectedResult:    sdk.OneDec(),
		},
		"nextTickSqrtPrice > sqrtPriceLimit -> nextTickSqrtPrice": {
			sqrtPriceLimit:    three,
			nextTickSqrtPrice: four,
			expectedResult:    four,
		},
		"nextTickSqrtPrice < sqrtPriceLimit -> sqrtPriceLimit": {
			sqrtPriceLimit:    five,
			nextTickSqrtPrice: two,
			expectedResult:    five,
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()

			sut := swapstrategy.New(true, tc.sqrtPriceLimit, suite.App.GetKey(types.ModuleName), sdk.ZeroDec())

			actualSqrtTargetPrice := sut.GetSqrtTargetPrice(tc.nextTickSqrtPrice)

			suite.Require().Equal(tc.expectedResult, actualSqrtTargetPrice)

		})
	}
}

func (suite *StrategyTestSuite) TestComputeSwapStepOutGivenIn_ZeroForOne() {
	tokenOutErrTolerance := osmomath.ErrTolerance{
		AdditiveTolerance: sdk.SmallestDec().MulInt64(200),
	}

	var (
		sqrtPriceCurrent = defaultSqrtPriceUpper
		sqrtPriceNext    = defaultSqrtPriceLower
		// Note, we take a ceiling here.
		defaultAmountZero = defaultAmountZero.Ceil()

		// liquidity * sqrtPriceCurrent / (liquidity + amount in * sqrtPriceCurrent)
		sqrtPriceTargetNotReached = sdk.MustNewDecFromStr("70.711006269285598091")
		// liquidity * (sqrtPriceCurrent - sqrtPriceNext)
		amountOneTargetNotReached = sdk.MustNewDecFromStr("5238179488.140735222342785879")
	)

	tests := map[string]struct {
		sqrtPriceCurrent      sdk.Dec
		sqrtPriceTarget       sdk.Dec
		liquidity             sdk.Dec
		amountZeroInRemaining sdk.Dec
		swapFee               sdk.Dec

		expectedSqrtPriceNext        sdk.Dec
		expectedNewAmountRemainingIn sdk.Dec
		expectedAmountOut            sdk.Dec
		expectedFeeChargeTotal       sdk.Dec

		expectError error
	}{
		"1: no fee - reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountZeroInRemaining: defaultAmountZero,
			swapFee:               sdk.ZeroDec(),

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountZero,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent)
			expectedAmountOut:      defaultAmountOne,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"2: no fee - do not reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountZeroInRemaining: defaultAmountZero.Sub(sdk.NewDec(100)),
			swapFee:               sdk.ZeroDec(),

			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountZero.Sub(sdk.NewDec(100)),

			expectedAmountOut:      amountOneTargetNotReached,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"3: 3% fee - reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountZeroInRemaining: defaultAmountZero.Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:               defaultFee,

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountZero,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent)
			expectedAmountOut:      defaultAmountOne,
			expectedFeeChargeTotal: defaultAmountZero.Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
		"4: 3% fee - do not reach target": {
			sqrtPriceCurrent:      sqrtPriceCurrent,
			sqrtPriceTarget:       sqrtPriceNext,
			liquidity:             defaultLiquidity,
			amountZeroInRemaining: defaultAmountZero.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:               defaultFee,

			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountZero.Sub(sdk.NewDec(100)),
			expectedAmountOut:            amountOneTargetNotReached,
			expectedFeeChargeTotal:       defaultAmountZero.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()
			strategy := swapstrategy.New(true, types.MaxSqrtRatio, suite.App.GetKey(types.ModuleName), tc.swapFee)

			sqrtPriceNext, newAmountRemainingOneIn, amountZeroOut, feeChargeTotal := strategy.ComputeSwapStepOutGivenIn(tc.sqrtPriceCurrent, tc.sqrtPriceTarget, tc.liquidity, tc.amountZeroInRemaining)

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

func (suite *StrategyTestSuite) TestComputeSwapStepInGivenOut_ZeroForOne() {
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
		sqrtPriceCurrent       sdk.Dec
		sqrtPriceTarget        sdk.Dec
		liquidity              sdk.Dec
		amountZeroOutRemaining sdk.Dec
		swapFee                sdk.Dec

		expectedSqrtPriceNext        sdk.Dec
		expectedNewAmountRemainingIn sdk.Dec
		expectedAmountOutOne         sdk.Dec
		expectedFeeChargeTotal       sdk.Dec

		expectError error
	}{
		"1: no fee - reach target": {
			sqrtPriceCurrent:       sqrtPriceCurrent,
			sqrtPriceTarget:        sqrtPriceNext,
			liquidity:              defaultLiquidity,
			amountZeroOutRemaining: defaultAmountZero,
			swapFee:                sdk.ZeroDec(),

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountZero,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
			expectedAmountOutOne:   defaultAmountOne,
			expectedFeeChargeTotal: sdk.ZeroDec(),
		},
		"2: no fee - do not reach target": {
			sqrtPriceCurrent:       sqrtPriceCurrent,
			sqrtPriceTarget:        sqrtPriceNext,
			liquidity:              defaultLiquidity,
			amountZeroOutRemaining: defaultAmountOne.Sub(sdk.NewDec(100)),
			swapFee:                sdk.ZeroDec(),

			// sqrt_price_current + token_in / liquidity
			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountOne.Sub(sdk.NewDec(100)),
			expectedAmountOutOne:         amountZeroTargetNotReached,
			expectedFeeChargeTotal:       sdk.ZeroDec(),
		},
		"3: 3% fee - reach target": {
			sqrtPriceCurrent:       sqrtPriceCurrent,
			sqrtPriceTarget:        sqrtPriceNext,
			liquidity:              defaultLiquidity,
			amountZeroOutRemaining: defaultAmountOne.Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:                defaultFee,

			expectedSqrtPriceNext:        sqrtPriceNext,
			expectedNewAmountRemainingIn: defaultAmountOne,
			// liquidity * (sqrtPriceNext - sqrtPriceCurrent) / (sqrtPriceNext * sqrtPriceCurrent)
			expectedAmountOutOne:   defaultAmountZero,
			expectedFeeChargeTotal: defaultAmountOne.Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
		"4: 3% fee - do not reach target": {
			sqrtPriceCurrent:       sqrtPriceCurrent,
			sqrtPriceTarget:        sqrtPriceNext,
			liquidity:              defaultLiquidity,
			amountZeroOutRemaining: defaultAmountOne.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)),
			swapFee:                defaultFee,

			expectedSqrtPriceNext:        sqrtPriceTargetNotReached,
			expectedNewAmountRemainingIn: defaultAmountOne.Sub(sdk.NewDec(100)),
			expectedAmountOutOne:         amountZeroTargetNotReached,
			expectedFeeChargeTotal:       defaultAmountOne.Sub(sdk.NewDec(100)).Quo(sdk.OneDec().Sub(defaultFee)).Mul(defaultFee),
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()
			strategy := swapstrategy.New(true, types.MaxSqrtRatio, suite.App.GetKey(types.ModuleName), tc.swapFee)

			sqrtPriceNext, newAmountRemainingZeroOut, amountOneIn, feeChargeTotal := strategy.ComputeSwapStepOutGivenIn(tc.sqrtPriceCurrent, tc.sqrtPriceTarget, tc.liquidity, tc.amountZeroOutRemaining)

			suite.Require().Equal(tc.expectedSqrtPriceNext, sqrtPriceNext)
			suite.Require().Equal(tc.expectedNewAmountRemainingIn, newAmountRemainingZeroOut)
			suite.Require().Equal(0,
				tokenOutErrTolerance.CompareBigDec(
					osmomath.BigDecFromSDKDec(tc.expectedAmountOutOne),
					osmomath.BigDecFromSDKDec(amountOneIn),
				),
				fmt.Sprintf("expected (%s), actual (%s)", tc.expectedAmountOutOne, amountOneIn))
			suite.Require().Equal(tc.expectedFeeChargeTotal, feeChargeTotal)
		})
	}
}
