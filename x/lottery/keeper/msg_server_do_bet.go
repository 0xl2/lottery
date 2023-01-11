package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

func (k msgServer) DoBet(goCtx context.Context, msg *types.MsgDoBet) (*types.MsgDoBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check user has bet already
	betted := k.Keeper.CheckBet(ctx, msg.Creator)
	if betted {
		panic("You can bet only 1 time")
	}

	// check user has sufficient balance
	hasBalance := k.Keeper.CheckBalance(ctx, msg.Creator)
	if !hasBalance {
		panic("Insufficient balance")
	}

	// check bet amount is in valid range
	validRange := k.Keeper.CheckRange(ctx, msg.BetAmount)
	if !validRange {
		panic("Invalid amount")
	}

	// save bet info
	betInfo, found := k.Keeper.GetBetInfo(ctx)
	if !found {
		panic("BetInfo not found")
	}

	storedBet := types.StoredBet{
		Index: string(rune(betInfo.BetId)),
		UserAddr: msg.Creator,
		BetAmount: msg.BetAmount,
	}
	k.Keeper.SetStoredBet(ctx, storedBet)

	// update betID
	betInfo.BetId++

	// update lottery info
	lotteryInfo, found := k.Keeper.GetLotteryInfo(ctx)
	if !found {
		panic("LotteryInfo not found")
	}
	lotteryInfo.NextOrder++

	return &types.MsgDoBetResponse{}, nil
}
