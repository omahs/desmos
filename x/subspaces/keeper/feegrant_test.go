package keeper_test

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SaveUserGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		grantee    string
		allowance  feegrant.FeeAllowanceI
		check      func(ctx sdk.Context)
	}{
		{
			name:       "grant is saved properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, err := suite.k.GetUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)

				// check if grant is set properly
				expected, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)

				// check if account is set properly
				addr, err := sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
		{
			name: "existing grant is overridden properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, err := suite.k.GetUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)

				// check if grant is set properly
				expected, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)

				// check if account is set properly
				addr, err := sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			grant, err := types.NewUserGrant(tc.subspaceID, tc.granter, tc.grantee, tc.allowance)
			suite.Require().NoError(err)
			suite.k.SaveUserGrant(ctx, grant)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasUserGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		grantee    string
		expResult  bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expResult:  false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			suite.Require().Equal(tc.expResult, suite.k.HasUserGrant(ctx, tc.subspaceID, tc.granter, tc.grantee))
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetUserGrant() {
	allowance, err := codectypes.NewAnyWithValue(&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
	suite.Require().NoError(err)

	invalidAllowance, err := codectypes.NewAnyWithValue(&types.MsgCreateSubspace{})
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		grantee    string
		shouldErr  bool
		expResult  types.UserGrant
	}{
		{
			name:       "non-existing grant returns error",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			shouldErr:  true,
		},
		{
			name: "invalid stored grant returns error",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, types.UserGrant{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  invalidAllowance,
				})
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			shouldErr:  true,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			shouldErr:  false,
			expResult: types.UserGrant{
				SubspaceID: 1,
				Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				Allowance:  allowance,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			grant, err := suite.k.GetUserGrant(ctx, tc.subspaceID, tc.granter, tc.grantee)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResult, grant)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteUserGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		grantee    string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing grant is deleted properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteUserGrant(ctx, tc.subspaceID, tc.granter, tc.grantee)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveGroupGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		allowance  feegrant.FeeAllowanceI
		check      func(ctx sdk.Context)
	}{
		{
			name:       "grant is saved properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, err := suite.k.GetGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1)
				suite.Require().NoError(err)

				// check if grant is set properly
				expected, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)
			},
		},
		{
			name: "existing group grant is overridden properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, err := suite.k.GetGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1)
				suite.Require().NoError(err)

				// check if grant is set properly
				expected, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			grant, err := types.NewGroupGrant(tc.subspaceID, tc.granter, tc.groupID, tc.allowance)
			suite.Require().NoError(err)
			suite.k.SaveGroupGrant(ctx, grant)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasGroupGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		expResult  bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expResult:  false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				grant, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveGroupGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			suite.Require().Equal(tc.expResult, suite.k.HasGroupGrant(ctx, tc.subspaceID, tc.granter, tc.groupID))
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetGroupGrant() {
	allowance, err := codectypes.NewAnyWithValue(&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
	suite.Require().NoError(err)

	invalidAllowance, err := codectypes.NewAnyWithValue(&types.MsgCreateSubspace{})
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		shouldErr  bool
		expResult  types.GroupGrant
	}{
		{
			name:       "non-existing grant returns error",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			shouldErr:  true,
		},
		{
			name: "invalid stored grant returns error",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(err)
				suite.k.SaveGroupGrant(ctx, types.GroupGrant{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					GroupID:    1,
					Allowance:  invalidAllowance,
				})
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			shouldErr:  true,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				grant, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveGroupGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			shouldErr:  false,
			expResult: types.GroupGrant{
				SubspaceID: 1,
				Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				GroupID:    1,
				Allowance:  allowance,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			grant, err := suite.k.GetGroupGrant(ctx, tc.subspaceID, tc.granter, tc.groupID)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResult, grant)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteGroupGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing grant is deleted properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1))
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteGroupGrant(ctx, tc.subspaceID, tc.granter, tc.groupID)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
