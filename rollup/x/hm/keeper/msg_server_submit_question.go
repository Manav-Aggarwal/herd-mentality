package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"hm/x/hm/types"
)

func (k msgServer) SubmitQuestion(goCtx context.Context, msg *types.MsgSubmitQuestion) (*types.MsgSubmitQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitQuestionResponse{}, nil
}
