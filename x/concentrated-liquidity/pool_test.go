package concentrated_liquidity_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	cl "github.com/osmosis-labs/osmosis/v12/x/concentrated-liquidity"
	cltypes "github.com/osmosis-labs/osmosis/v12/x/concentrated-liquidity/types"
)

func (s *KeeperTestSuite) TestCalcOutAmtGivenIn() {
	currPrice := sdk.NewDec(5000)
	currSqrtPrice, err := currPrice.ApproxSqrt() // 70.710678118654752440
	s.Require().NoError(err)
	currTick := cl.PriceToTick(currPrice) // 85176
	lowerPrice := sdk.NewDec(4545)
	s.Require().NoError(err)
	lowerTick := cl.PriceToTick(lowerPrice) // 84222
	upperPrice := sdk.NewDec(5500)
	s.Require().NoError(err)
	upperTick := cl.PriceToTick(upperPrice) // 86129

	defaultAmt0 := sdk.NewInt(1000000)
	defaultAmt1 := sdk.NewInt(5000000000)

	swapFee := sdk.ZeroDec()

	tests := map[string]struct {
		positionAmount0  sdk.Int
		positionAmount1  sdk.Int
		addPositions     func(ctx sdk.Context, poolId uint64)
		tokenIn          sdk.Coin
		tokenOutDenom    string
		priceLimit       sdk.Dec
		expectedTokenIn  sdk.Coin
		expectedTokenOut sdk.Coin
		expectedTick     sdk.Int
		newLowerPrice    sdk.Dec
		newUpperPrice    sdk.Dec
		poolLiqAmount0   sdk.Int
		poolLiqAmount1   sdk.Int
	}{
		//  One price range
		//
		//          5000
		//  4545 -----|----- 5500
		"single position within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5004),
			// params
			// liquidity: 		 1517818840.967515822610790519
			// sqrtPriceNext:    70.738349405152439867 which is 5003.914076565430543175
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  42000000.0000 rounded up https://www.wolframalpha.com/input?i=1517818840.967515822610790519+*+%2870.738349405152439867+-+70.710678118654752440%29
			// expectedTokenOut: 8396.714105 rounded down https://www.wolframalpha.com/input?i=%281517818840.967515822610790519+*+%2870.738349405152439867+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.738349405152439867%29
			// expectedTick: 	 85184.0 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C5003.914076565430543175%5D
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8396)),
			expectedTick:     sdk.NewInt(85184),
		},
		"single position within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4993),
			// params
			// liquidity: 		 1517818840.967515822610790519
			// sqrtPriceNext:    70.666662070529219856 which is 4993.777128190373086350
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13369.9999 rounded up https://www.wolframalpha.com/input?i=%281517818840.967515822610790519+*+%2870.710678118654752440+-+70.666662070529219856+%29%29+%2F+%2870.666662070529219856+*+70.710678118654752440%29
			// expectedTokenOut: 66808387.149 rounded down https://www.wolframalpha.com/input?i=1517818840.967515822610790519+*+%2870.710678118654752440+-+70.666662070529219856%29
			// expectedTick: 	 85163.7 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C4993.777128190373086350%5D
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66808387)),
			expectedTick:     sdk.NewInt(85163),
		},
		//  Two equal price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//  4545 -----|----- 5500
		"two positions within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5002),
			// params
			// liquidity: 		 3035637681.935031645221581038
			// sqrtPriceNext:    70.724513761903596153 which is 5001.956846857691162236
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  41999999.999 rounded up https://www.wolframalpha.com/input?i=3035637681.935031645221581038+*+%2870.724513761903596153+-+70.710678118654752440%29
			// expectedTokenOut: 8398.3567 rounded down https://www.wolframalpha.com/input?i=%283035637681.935031645221581038+*+%2870.724513761903596153+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+70.724513761903596153%29
			// expectedTick:     85180.1 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C5003.914076565430543175%5D
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8398)),
			expectedTick:     sdk.NewInt(85180),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		"two positions within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom: "usdc",
			priceLimit:    sdk.NewDec(4996),
			// we expect to put .01337 eth in and in return get 66.76 eth back
			// TODO: look into why we are returning 66.78 instead of 66.76 like the inverse of this test above
			// sure, the above test only has 1 position while this has two positions, but shouldn't that effect the tokenIn as well?
			// params
			// liquidity: 		 3035637681.935031645221581038
			// sqrtPriceNext:    70.688663242671855280 which is 4996.887111035867053835
			// sqrtPriceCurrent: 70.710678118654752440 which is 5000
			// expectedTokenIn:  13369.9999 rounded up https://www.wolframalpha.com/input?i=%283035637681.935031645221581038+*+%2870.710678118654752440+-+70.688663242671855280+%29%29+%2F+%2870.688663242671855280+*+70.710678118654752440%29
			// expectedTokenOut: 66829187.096 rounded down https://www.wolframalpha.com/input?i=3035637681.935031645221581038+*+%2870.710678118654752440+-+70.688663242671855280%29
			// expectedTick: 	 85170.00 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C4996.887111035867053835%5D
			expectedTokenIn:  sdk.NewCoin("eth", sdk.NewInt(13370)),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66829187)),
			expectedTick:     sdk.NewInt(85169), // TODO: should be 85170, is 85169 due to log precision
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		//  Consecutive price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//             5500 ----------- 6250
		//
		"two positions with consecutive price ranges": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (1st):  1517818840.967515822610790519
				// sqrtPriceNext:    74.160724590951092256 which is 5499.813071854898049815
				// sqrtPriceCurrent: 70.710678118654752440 which is 5000
				// expectedTokenIn:  5236545537.865 rounded up https://www.wolframalpha.com/input?i=1517818840.967515822610790519+*+%2874.160724590951092256+-+70.710678118654752440%29
				// expectedTokenOut: 998587.023 rounded down https://www.wolframalpha.com/input?i=%281517818840.967515822610790519+*+%2874.161984870956629487+-+70.710678118654752440+%29%29+%2F+%2870.710678118654752440+*+74.161984870956629487%29
				// expectedTick:     86129.0 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C5499.813071854898049815%5D

				// create second position parameters
				newLowerPrice := sdk.NewDec(5501)
				s.Require().NoError(err)
				newLowerTick := cl.PriceToTick(newLowerPrice) // 84222
				newUpperPrice := sdk.NewDec(6250)
				s.Require().NoError(err)
				newUpperTick := cl.PriceToTick(newUpperPrice) // 87407

				// add position two with the new price range above
				_, _, _, err := s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[2], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
				// params
				// liquidity (2nd):  1200046517.432645168444000000
				// sqrtPriceNext:    78.137532175162292735 which is 6105.473934424522538231
				// sqrtPriceCurrent: 74.168140663410187419 which is 5500.913089467399755950
				// expectedTokenIn:  4763454460.005 rounded up https://www.wolframalpha.com/input?i=1200046517.432645168444000000+*+%2878.137532175162292735+-+74.168140663410187419%29
				// expectedTokenOut: 821949.1205 rounded down https://www.wolframalpha.com/input?i=%281200046517.432645168444000000+*+%2878.137532175162292735+-+74.168140663410187419+%29%29+%2F+%2874.168140663410187419+*+78.137532175162292735%29
				// expectedTick:     87173.8 rounded down https://www.wolframalpha.com/input?i2d=true&i=Log%5B1.0001%2C6105.473934424522538231%5D
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(6106),
			// expectedTokenIn:  5236545538 + 4763454461 = 9999999999 rounded up = 1000000000 = 100000 usdc
			// expectedTokenOut: 998587 + 821949 = 1820536 = 1.820536 eth
			expectedTokenIn:  sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1820536)),
			expectedTick:     sdk.NewInt(87173),
			newLowerPrice:    sdk.NewDec(5501),
			newUpperPrice:    sdk.NewDec(6250),
		},
	}

	for name, test := range tests {
		s.Run(name, func() {
			s.Setup()
			// create pool
			pool, err := s.App.ConcentratedLiquidityKeeper.CreateNewConcentratedLiquidityPool(s.Ctx, 1, "eth", "usdc", currSqrtPrice, currTick)
			s.Require().NoError(err)

			// add positions
			test.addPositions(s.Ctx, pool.Id)

			tokenIn, tokenOut, updatedTick, updatedLiquidity, _, err := s.App.ConcentratedLiquidityKeeper.CalcOutAmtGivenIn(
				s.Ctx,
				test.tokenIn, test.tokenOutDenom,
				swapFee, test.priceLimit, pool.Id)
			s.Require().NoError(err)

			s.Require().Equal(test.expectedTokenIn.String(), tokenIn.String())
			s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())
			s.Require().Equal(test.expectedTick.String(), updatedTick.String())

			if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
				test.newLowerPrice = lowerPrice
				test.newUpperPrice = upperPrice
			}

			newLowerTick := cl.PriceToTick(test.newLowerPrice)
			newUpperTick := cl.PriceToTick(test.newUpperPrice)

			lowerSqrtPrice, err := cl.TickToSqrtPrice(newLowerTick)
			s.Require().NoError(err)
			upperSqrtPrice, err := cl.TickToSqrtPrice(newUpperTick)
			s.Require().NoError(err)

			if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
				test.poolLiqAmount0 = defaultAmt0
				test.poolLiqAmount1 = defaultAmt1
			}

			expectedLiquidity := cl.GetLiquidityFromAmounts(currSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
			s.Require().Equal(expectedLiquidity.TruncateInt(), updatedLiquidity.TruncateInt())

			//poolNew := s.App.ConcentratedLiquidityKeeper.GetPoolbyId(s.Ctx, pool.Id)
			//fmt.Printf("%v pool price \n", poolNew.CurrentSqrtPrice.Power(2))
		})

	}
}

func (s *KeeperTestSuite) TestSwapOutAmtGivenIn() {
	currPrice := sdk.NewDec(5000)
	currSqrtPrice, err := currPrice.ApproxSqrt() // 70.710678118654752440
	s.Require().NoError(err)
	currTick := cl.PriceToTick(currPrice) // 85176
	lowerPrice := sdk.NewDec(4545)
	s.Require().NoError(err)
	lowerTick := cl.PriceToTick(lowerPrice) // 84222
	upperPrice := sdk.NewDec(5500)
	s.Require().NoError(err)
	upperTick := cl.PriceToTick(upperPrice) // 86129

	defaultAmt0 := sdk.NewInt(1000000)
	defaultAmt1 := sdk.NewInt(5000000000)

	swapFee := sdk.ZeroDec()

	tests := map[string]struct {
		positionAmount0  sdk.Int
		positionAmount1  sdk.Int
		addPositions     func(ctx sdk.Context, poolId uint64)
		tokenIn          sdk.Coin
		tokenOutDenom    string
		priceLimit       sdk.Dec
		expectedTokenOut sdk.Coin
		expectedTick     sdk.Int
		newLowerPrice    sdk.Dec
		newUpperPrice    sdk.Dec
		poolLiqAmount0   sdk.Int
		poolLiqAmount1   sdk.Int
	}{
		//  One price range
		//
		//          5000
		//  4545 -----|----- 5500
		"single position within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5004),
			// we expect to put 42 usdc in and in return get .008398 eth back
			// due to truncations and precision loss, we actually put in 41.99 usdc and in return get .008396 eth back
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8396)),
			expectedTick:     sdk.NewInt(85184),
		},
		"single position within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:          sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom:    "usdc",
			priceLimit:       sdk.NewDec(4993),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66808387)),
			expectedTick:     sdk.NewInt(85163),
		},
		//  Two equal price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//  4545 -----|----- 5500
		"two positions within one tick: usdc -> eth": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:       sdk.NewCoin("usdc", sdk.NewInt(42000000)),
			tokenOutDenom: "eth",
			priceLimit:    sdk.NewDec(5002),
			// we expect to put 42 usdc in and in return get .008398 eth back
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(8398)),
			expectedTick:     sdk.NewInt(85180),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		"two positions within one tick: eth -> usdc": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)

				// add second position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[1], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:          sdk.NewCoin("eth", sdk.NewInt(13370)),
			tokenOutDenom:    "usdc",
			priceLimit:       sdk.NewDec(4996),
			expectedTokenOut: sdk.NewCoin("usdc", sdk.NewInt(66829187)),
			expectedTick:     sdk.NewInt(85169),
			// two positions with same liquidity entered
			poolLiqAmount0: sdk.NewInt(1000000).MulRaw(2),
			poolLiqAmount1: sdk.NewInt(5000000000).MulRaw(2),
		},
		//  Consecutive price ranges
		//
		//          5000
		//  4545 -----|----- 5500
		//             5500 ----------- 6250
		//
		"two positions with consecutive price ranges": {
			addPositions: func(ctx sdk.Context, poolId uint64) {
				// add first position
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[0], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), lowerTick.Int64(), upperTick.Int64())
				s.Require().NoError(err)

				// create second position parameters
				newLowerPrice := sdk.NewDec(5501)
				s.Require().NoError(err)
				newLowerTick := cl.PriceToTick(newLowerPrice) // 84222
				newUpperPrice := sdk.NewDec(6250)
				s.Require().NoError(err)
				newUpperTick := cl.PriceToTick(newUpperPrice) // 87407

				// add position two with the new price range above
				_, _, _, err = s.App.ConcentratedLiquidityKeeper.CreatePosition(ctx, poolId, s.TestAccs[2], defaultAmt0, defaultAmt1, sdk.ZeroInt(), sdk.ZeroInt(), newLowerTick.Int64(), newUpperTick.Int64())
				s.Require().NoError(err)
			},
			tokenIn:          sdk.NewCoin("usdc", sdk.NewInt(10000000000)),
			tokenOutDenom:    "eth",
			priceLimit:       sdk.NewDec(6106),
			expectedTokenOut: sdk.NewCoin("eth", sdk.NewInt(1820536)),
			expectedTick:     sdk.NewInt(87173),
			newLowerPrice:    sdk.NewDec(5501),
			newUpperPrice:    sdk.NewDec(6250),
		},
	}

	for name, test := range tests {
		s.Run(name, func() {
			s.Setup()
			// create pool
			pool, err := s.App.ConcentratedLiquidityKeeper.CreateNewConcentratedLiquidityPool(s.Ctx, 1, "eth", "usdc", currSqrtPrice, currTick)
			s.Require().NoError(err)

			// add positions
			test.addPositions(s.Ctx, pool.Id)

			// execute internal swap function
			tokenOut, err := s.App.ConcentratedLiquidityKeeper.SwapOutAmtGivenIn(
				s.Ctx,
				test.tokenIn, test.tokenOutDenom,
				swapFee, test.priceLimit, pool.Id)
			s.Require().NoError(err)

			pool = s.App.ConcentratedLiquidityKeeper.GetPoolbyId(s.Ctx, pool.Id)
			s.Require().NoError(err)

			// check that we produced the same token out from the swap function that we expected
			s.Require().Equal(test.expectedTokenOut.String(), tokenOut.String())

			// check that the pool's current tick was updated correctly
			s.Require().Equal(test.expectedTick.String(), pool.CurrentTick.String())

			// the following is needed to get the expected liquidity to later compare to what the pool was updated to
			if test.newLowerPrice.IsNil() && test.newUpperPrice.IsNil() {
				test.newLowerPrice = lowerPrice
				test.newUpperPrice = upperPrice
			}

			newLowerTick := cl.PriceToTick(test.newLowerPrice)
			newUpperTick := cl.PriceToTick(test.newUpperPrice)

			lowerSqrtPrice, err := cl.TickToSqrtPrice(newLowerTick)
			s.Require().NoError(err)
			upperSqrtPrice, err := cl.TickToSqrtPrice(newUpperTick)
			s.Require().NoError(err)

			if test.poolLiqAmount0.IsNil() && test.poolLiqAmount1.IsNil() {
				test.poolLiqAmount0 = defaultAmt0
				test.poolLiqAmount1 = defaultAmt1
			}

			expectedLiquidity := cl.GetLiquidityFromAmounts(currSqrtPrice, lowerSqrtPrice, upperSqrtPrice, test.poolLiqAmount0, test.poolLiqAmount1)
			// check that the pools liquidity was updated correctly
			s.Require().Equal(expectedLiquidity.TruncateInt(), pool.Liquidity.TruncateInt())

			// TODO: need to figure out a good way to test that the currentSqrtPrice that the pool is set to makes sense
			// right now we calculate this value through iterations, so unsure how to do this here / if its needed
		})

	}
}

func (s *KeeperTestSuite) TestOrderInitialPoolDenoms() {
	denom0, denom1, err := cltypes.OrderInitialPoolDenoms("axel", "osmo")
	s.Require().NoError(err)
	s.Require().Equal(denom0, "axel")
	s.Require().Equal(denom1, "osmo")

	denom0, denom1, err = cltypes.OrderInitialPoolDenoms("usdc", "eth")
	s.Require().NoError(err)
	s.Require().Equal(denom0, "eth")
	s.Require().Equal(denom1, "usdc")

	denom0, denom1, err = cltypes.OrderInitialPoolDenoms("usdc", "usdc")
	s.Require().Error(err)

}

func (suite *KeeperTestSuite) TestPriceToTick() {
	testCases := []struct {
		name         string
		price        sdk.Dec
		tickExpected string
	}{
		{
			"happy path",
			sdk.NewDec(5000),
			"85176",
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			tick := cl.PriceToTick(tc.price)
			suite.Require().Equal(tc.tickExpected, tick.String())
		})
	}
}