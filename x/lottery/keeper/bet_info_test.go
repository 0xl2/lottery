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

func createTestBetInfo(keeper *keeper.Keeper, ctx sdk.Context) types.BetInfo {
	item := types.BetInfo{}
	keeper.SetBetInfo(ctx, item)
	return item
}

func TestBetInfoGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestBetInfo(keeper, ctx)
	rst, found := keeper.GetBetInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestBetInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestBetInfo(keeper, ctx)
	keeper.RemoveBetInfo(ctx)
	_, found := keeper.GetBetInfo(ctx)
	require.False(t, found)
}
