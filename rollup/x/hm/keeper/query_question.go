package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hm/x/hm/types"
)

func (k Keeper) QuestionAll(ctx context.Context, req *types.QueryAllQuestionRequest) (*types.QueryAllQuestionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var questions []types.Question

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	questionStore := prefix.NewStore(store, types.KeyPrefix(types.QuestionKeyPrefix))

	pageRes, err := query.Paginate(questionStore, req.Pagination, func(key []byte, value []byte) error {
		var question types.Question
		if err := k.cdc.Unmarshal(value, &question); err != nil {
			return err
		}

		questions = append(questions, question)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllQuestionResponse{Question: questions, Pagination: pageRes}, nil
}

func (k Keeper) Question(ctx context.Context, req *types.QueryGetQuestionRequest) (*types.QueryGetQuestionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetQuestion(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetQuestionResponse{Question: val}, nil
}
