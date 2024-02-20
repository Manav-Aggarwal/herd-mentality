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

func (k Keeper) AnswerAll(ctx context.Context, req *types.QueryAllAnswerRequest) (*types.QueryAllAnswerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var answers []types.Answer

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	answerStore := prefix.NewStore(store, types.KeyPrefix(types.AnswerKeyPrefix))

	pageRes, err := query.Paginate(answerStore, req.Pagination, func(key []byte, value []byte) error {
		var answer types.Answer
		if err := k.cdc.Unmarshal(value, &answer); err != nil {
			return err
		}

		answers = append(answers, answer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAnswerResponse{Answer: answers, Pagination: pageRes}, nil
}

func (k Keeper) Answer(ctx context.Context, req *types.QueryGetAnswerRequest) (*types.QueryGetAnswerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAnswer(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAnswerResponse{Answer: val}, nil
}
