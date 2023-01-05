package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

func (k msgServer) DoBet(goCtx context.Context, msg *types.MsgDoBet) (*types.MsgDoBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check user has sufficient balance
	hasBalance := k.Keeper.CheckBalance(ctx, msg.Creator)
	if !hasBalance {
		panic("Insufficient balance")
	}

	// // Get betInfo
	// betInfo, found := k.Keeper.GetBetInfo(ctx)
	// if !found {
	// 	panic("BetInfo not found")
	// }
	// // Get lotteryInfo
	// lotteryInfo, found := k.Keeper.GetLotteryInfo(ctx)
	// if !found {
	// 	panic("LotteryInfo not found")
	// }

	return &types.MsgDoBetResponse{}, nil
}
