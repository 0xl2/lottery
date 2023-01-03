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

func createNStoredLottery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredLottery {
	items := make([]types.StoredLottery, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredLottery(ctx, items[i])
	}
	return items
}

func TestStoredLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredLottery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredLottery(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredLottery(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredLottery(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredLottery(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredLotteryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNStoredLottery(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredLottery(ctx)),
	)
}
