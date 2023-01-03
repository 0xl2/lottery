package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

// SetLotteryData set lotteryData in the store
func (k Keeper) SetLotteryData(ctx sdk.Context, lotteryData types.LotteryData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryDataKey))
	b := k.cdc.MustMarshal(&lotteryData)
	store.Set([]byte{0}, b)
}

// GetLotteryData returns lotteryData
func (k Keeper) GetLotteryData(ctx sdk.Context) (val types.LotteryData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryDataKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotteryData removes lotteryData from the store
func (k Keeper) RemoveLotteryData(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryDataKey))
	store.Delete([]byte{0})
}
