package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")

	ErrAlreadyBet = sdkerrors.Register(ModuleName, 1200, "You can bet only 1 time")
	ErrInsufficient = sdkerrors.Register(ModuleName, 1201, "Insufficient balance")
	ErrInvalidRange = sdkerrors.Register(ModuleName, 1202, "Invalid bet amount range")

	ErrLotteryData = sdkerrors.Register(ModuleName, 1300, "LotteryData not found")
	ErrLotteryInfo = sdkerrors.Register(ModuleName, 1301, "LotteryInfo not found")
	ErrLottery = sdkerrors.Register(ModuleName, 1302, "Lottery not found")
	ErrBetTx = sdkerrors.Register(ModuleName, 1303, "Bet tx not found")
	ErrBetInfo = sdkerrors.Register(ModuleName, 1304, "BetInfo not found")
)
