package lottery_test

import (
	"testing"

	keepertest "github.com/naruto0913/lottery/testutil/keeper"
	"github.com/naruto0913/lottery/testutil/nullify"
	"github.com/naruto0913/lottery/x/lottery"
	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LotteryInfo: &types.LotteryInfo{
			NextId: 52,
		},
		StoredLotteryList: []types.StoredLottery{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		StoredBetList: []types.StoredBet{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BetInfo: &types.BetInfo{
			BetId: 91,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.LotteryInfo, got.LotteryInfo)
	require.ElementsMatch(t, genesisState.StoredLotteryList, got.StoredLotteryList)
	require.ElementsMatch(t, genesisState.StoredBetList, got.StoredBetList)
	require.Equal(t, genesisState.BetInfo, got.BetInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
