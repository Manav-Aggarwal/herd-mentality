package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitAnswer{}

func NewMsgSubmitAnswer(creator string, qid string, player string, answer string) *MsgSubmitAnswer {
	return &MsgSubmitAnswer{
		Creator: creator,
		Qid:     qid,
		Player:  player,
		Answer:  answer,
	}
}

func (msg *MsgSubmitAnswer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
