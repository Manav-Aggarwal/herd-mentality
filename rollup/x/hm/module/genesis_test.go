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

		QuestionList: []types.Question{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		AnswerList: []types.Answer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HmKeeper(t)
	hm.InitGenesis(ctx, k, genesisState)
	got := hm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.QuestionList, got.QuestionList)
	require.ElementsMatch(t, genesisState.AnswerList, got.AnswerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
