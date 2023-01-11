package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

func (k msgServer) DoBet(goCtx context.Context, msg *types.MsgDoBet) (*types.MsgDoBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check user has bet already
	betted, err := k.Keeper.CheckBet(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if betted {
		return nil, types.ErrAlreadyBet
	}

	// check user has sufficient balance
	hasBalance, err := k.Keeper.CheckBalance(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if !hasBalance {
		return nil, types.ErrInsufficient
	}

	// check bet amount is in valid range
	validRange, err := k.Keeper.CheckRange(ctx, msg.BetAmount)
	if err != nil {
		return nil, err
	}
	if !validRange {
		return nil, types.ErrInvalidRange
	}

	respBet, err := k.Keeper.SaveBet(ctx, msg.Creator, msg.BetAmount)
	if err != nil {
		return nil, err
	}

	return respBet, nil
}
