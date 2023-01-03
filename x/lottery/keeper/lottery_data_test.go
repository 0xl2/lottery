package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/naruto0913/lottery/testutil/keeper"
	"github.com/naruto0913/lottery/testutil/nullify"
	"github.com/naruto0913/lottery/x/lottery/keeper"
	"github.com/naruto0913/lottery/x/lottery/types"
)

func createTestLotteryData(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryData {
	item := types.LotteryData{}
	keeper.SetLotteryData(ctx, item)
	return item
}

func TestLotteryDataGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLotteryData(keeper, ctx)
	rst, found := keeper.GetLotteryData(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLotteryDataRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLotteryData(keeper, ctx)
	keeper.RemoveLotteryData(ctx)
	_, found := keeper.GetLotteryData(ctx)
	require.False(t, found)
}
