package keeper

import (
	"context"

	"hm/x/hm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitQuestion(goCtx context.Context, msg *types.MsgSubmitQuestion) (*types.MsgSubmitQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Create the question
	var question = types.Question{
		Index: msg.Qid,
		Desc:  msg.Question,
	}
	k.SetQuestion(ctx, question)

	return &types.MsgSubmitQuestionResponse{}, nil
}
