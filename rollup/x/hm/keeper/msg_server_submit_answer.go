package keeper

import (
	"context"

	"hm/x/hm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitAnswer(goCtx context.Context, msg *types.MsgSubmitAnswer) (*types.MsgSubmitAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the question
	var answer = types.Answer{
		Index:  msg.Qid,
		Player: msg.Player,
		Desc:   msg.Answer,
	}
	k.SetAnswer(ctx, answer)

	return &types.MsgSubmitAnswerResponse{}, nil
}
