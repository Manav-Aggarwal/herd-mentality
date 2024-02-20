package hm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hm/testutil/keeper"
	"hm/testutil/nullify"
	"hm/x/hm/module"
	"hm/x/hm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HmKeeper(t)
	hm.InitGenesis(ctx, k, genesisState)
	got := hm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
