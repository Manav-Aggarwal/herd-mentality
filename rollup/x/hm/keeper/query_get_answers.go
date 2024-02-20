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

	return &types.QueryGetAnswersResponse{Answers: findMostOccurringStrings(answers), Pagination: pageRes}, nil
}

// FindMostOccurringItems returns a slice of the most frequently occurring Item instances based on the Value field.
func findMostOccurringStrings(answers []types.Answer) []types.Answer {
	// Map to store the count of each unique string value
	counts := make(map[string]int)
	// Map to store the last seen Item for each unique string value
	lastSeen := make(map[string]types.Answer)
	maxCount := 0

	// Count the occurrences of each unique string value
	for _, answer := range answers {
		value := answer.Desc
		counts[value]++
		lastSeen[value] = answer // Keep the last seen instance of each unique value
		if counts[value] > maxCount {
			maxCount = counts[value]
		}
	}

	// Find the Item(s) with the highest occurrence
	var mostOccurringItems []types.Answer
	for _, answer := range lastSeen {
		if counts[answer.Desc] == maxCount {
			mostOccurringItems = append(mostOccurringItems, answer)
		}
	}

	return mostOccurringItems
}
