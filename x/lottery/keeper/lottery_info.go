package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

// SetLotteryInfo set lotteryInfo in the store
func (k Keeper) SetLotteryInfo(ctx sdk.Context, lotteryInfo types.LotteryInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryInfoKey))
	b := k.cdc.MustMarshal(&lotteryInfo)
	store.Set([]byte{0}, b)
}

// GetLotteryInfo returns lotteryInfo
func (k Keeper) GetLotteryInfo(ctx sdk.Context) (val types.LotteryInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotteryInfo removes lotteryInfo from the store
func (k Keeper) RemoveLotteryInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryInfoKey))
	store.Delete([]byte{0})
}
