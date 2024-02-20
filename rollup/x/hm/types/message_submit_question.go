package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitQuestion{}

func NewMsgSubmitQuestion(creator string, qid string, question string) *MsgSubmitQuestion {
	return &MsgSubmitQuestion{
		Creator:  creator,
		Qid:      qid,
		Question: question,
	}
}

func (msg *MsgSubmitQuestion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
