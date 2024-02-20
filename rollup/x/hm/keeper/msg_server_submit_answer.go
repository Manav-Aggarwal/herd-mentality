package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"hm/x/hm/types"
)

func (k msgServer) SubmitAnswer(goCtx context.Context, msg *types.MsgSubmitAnswer) (*types.MsgSubmitAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitAnswerResponse{}, nil
}
