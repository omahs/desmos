package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/stretchr/testify/require"
)

func TestGroupGrantee_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		grantee   *types.GroupGrantee
		shouldErr bool
	}{
		{
			name:      "invalid group id returns error",
			grantee:   types.NewGroupGrantee(0),
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grantee:   types.NewGroupGrantee(1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grantee.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserGrantee_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		grantee   *types.UserGrantee
		shouldErr bool
	}{
		{
			name:      "invalid user address returns error",
			grantee:   types.NewUserGrantee(""),
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grantee:   types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grantee.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGrant_Validate(t *testing.T) {
	validUserGrantee := types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez")
	validgranteeAny, err := codectypes.NewAnyWithValue(types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"))
	require.NoError(t, err)

	userGrant, err := types.NewGrant(1, "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0", validUserGrantee, &feegrant.BasicAllowance{})
	require.NoError(t, err)

	invalidGranteeAny, err := codectypes.NewAnyWithValue(types.NewUserGrantee(""))
	require.NoError(t, err)

	invaliduserGranteeAny, err := codectypes.NewAnyWithValue(types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"))
	require.NoError(t, err)

	testCases := []struct {
		name      string
		grant     types.Grant
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			grant: types.Grant{
				SubspaceID: 0,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    &codectypes.Any{},
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "",
				Grantee:    &codectypes.Any{},
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    invalidGranteeAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "granter self-grant returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    invaliduserGranteeAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid allowance returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    validgranteeAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grant:     userGrant,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grant.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
