package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hm/testutil/keeper"
	"hm/testutil/nullify"
	"hm/x/hm/keeper"
	"hm/x/hm/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNQuestion(keeper keeper.Keeper, ctx context.Context, n int) []types.Question {
	items := make([]types.Question, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetQuestion(ctx, items[i])
	}
	return items
}

func TestQuestionGet(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNQuestion(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetQuestion(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestQuestionRemove(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNQuestion(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveQuestion(ctx,
			item.Index,
		)
		_, found := keeper.GetQuestion(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestQuestionGetAll(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNQuestion(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllQuestion(ctx)),
	)
}
