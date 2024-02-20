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

func createNAnswer(keeper keeper.Keeper, ctx context.Context, n int) []types.Answer {
	items := make([]types.Answer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAnswer(ctx, items[i])
	}
	return items
}

func TestAnswerGet(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNAnswer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAnswer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAnswerRemove(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNAnswer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAnswer(ctx,
			item.Index,
		)
		_, found := keeper.GetAnswer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAnswerGetAll(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	items := createNAnswer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAnswer(ctx)),
	)
}
