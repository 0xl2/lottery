package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/naruto0913/lottery/testutil/keeper"
	"github.com/naruto0913/lottery/testutil/nullify"
	"github.com/naruto0913/lottery/x/lottery/keeper"
	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredBet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredBet {
	items := make([]types.StoredBet, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredBet(ctx, items[i])
	}
	return items
}

func TestStoredBetGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBet(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredBet(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredBetRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBet(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredBet(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredBet(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredBetGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredBet(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredBet(ctx)),
	)
}
