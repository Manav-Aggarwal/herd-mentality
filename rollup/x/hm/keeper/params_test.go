package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hm/testutil/keeper"
	"hm/x/hm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HmKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
