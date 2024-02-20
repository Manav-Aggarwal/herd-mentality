package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "hm/testutil/keeper"
	"hm/testutil/nullify"
	"hm/x/hm/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestQuestionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	msgs := createNQuestion(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetQuestionRequest
		response *types.QueryGetQuestionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetQuestionRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetQuestionResponse{Question: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetQuestionRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetQuestionResponse{Question: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetQuestionRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Question(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestQuestionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.HmKeeper(t)
	msgs := createNQuestion(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllQuestionRequest {
		return &types.QueryAllQuestionRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.QuestionAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Question), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Question),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.QuestionAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Question), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Question),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.QuestionAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Question),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.QuestionAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
