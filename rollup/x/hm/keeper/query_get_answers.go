package keeper

import (
	"context"

	"hm/x/hm/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAnswers(goCtx context.Context, req *types.QueryGetAnswersRequest) (*types.QueryGetAnswersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query

	_ = ctx

	var answers []types.Answer

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	answerStore := prefix.NewStore(store, types.KeyPrefix(types.AnswerKeyPrefix))

	pageRes, err := query.Paginate(answerStore, req.Pagination, func(key []byte, value []byte) error {
		var answer types.Answer
		if err := k.cdc.Unmarshal(value, &answer); err != nil {
			return err
		}

		if req.Qid == answer.Index {
			answers = append(answers, answer)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAnswersResponse{Answers: answers, Pagination: pageRes}, nil
}
