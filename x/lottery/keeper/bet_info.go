package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

// SetBetInfo set betInfo in the store
func (k Keeper) SetBetInfo(ctx sdk.Context, betInfo types.BetInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetInfoKey))
	b := k.cdc.MustMarshal(&betInfo)
	store.Set([]byte{0}, b)
}

// GetBetInfo returns betInfo
func (k Keeper) GetBetInfo(ctx sdk.Context) (val types.BetInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBetInfo removes betInfo from the store
func (k Keeper) RemoveBetInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetInfoKey))
	store.Delete([]byte{0})
}
