package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDoBet = "do_bet"

var _ sdk.Msg = &MsgDoBet{}

func NewMsgDoBet(creator string, betAmount uint64) *MsgDoBet {
	return &MsgDoBet{
		Creator:   creator,
		BetAmount: betAmount,
	}
}

func (msg *MsgDoBet) Route() string {
	return RouterKey
}

func (msg *MsgDoBet) Type() string {
	return TypeMsgDoBet
}

func (msg *MsgDoBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDoBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDoBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
