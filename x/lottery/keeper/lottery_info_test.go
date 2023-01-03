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

func createTestLotteryInfo(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryInfo {
	item := types.LotteryInfo{}
	keeper.SetLotteryInfo(ctx, item)
	return item
}

func TestLotteryInfoGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLotteryInfo(keeper, ctx)
	rst, found := keeper.GetLotteryInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLotteryInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLotteryInfo(keeper, ctx)
	keeper.RemoveLotteryInfo(ctx)
	_, found := keeper.GetLotteryInfo(ctx)
	require.False(t, found)
}
