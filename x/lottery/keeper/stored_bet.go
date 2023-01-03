package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/types"
)

// SetStoredBet set a specific storedBet in the store from its index
func (k Keeper) SetStoredBet(ctx sdk.Context, storedBet types.StoredBet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBetKeyPrefix))
	b := k.cdc.MustMarshal(&storedBet)
	store.Set(types.StoredBetKey(
		storedBet.Index,
	), b)
}

// GetStoredBet returns a storedBet from its index
func (k Keeper) GetStoredBet(
	ctx sdk.Context,
	index string,

) (val types.StoredBet, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBetKeyPrefix))

	b := store.Get(types.StoredBetKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredBet removes a storedBet from the store
func (k Keeper) RemoveStoredBet(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBetKeyPrefix))
	store.Delete(types.StoredBetKey(
		index,
	))
}

// GetAllStoredBet returns all storedBet
func (k Keeper) GetAllStoredBet(ctx sdk.Context) (list []types.StoredBet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredBetKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredBet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
