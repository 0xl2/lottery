package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/naruto0913/lottery/x/lottery/keeper"
	"github.com/naruto0913/lottery/x/lottery/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.LotteryInfo != nil {
		k.SetLotteryInfo(ctx, *genState.LotteryInfo)
	}
	// Set all the storedLottery
	for _, elem := range genState.StoredLotteryList {
		k.SetStoredLottery(ctx, elem)
	}
	// Set all the storedBet
	for _, elem := range genState.StoredBetList {
		k.SetStoredBet(ctx, elem)
	}
	// Set if defined
	if genState.BetInfo != nil {
		k.SetBetInfo(ctx, *genState.BetInfo)
	}
	// Set if defined
	if genState.LotteryData != nil {
		k.SetLotteryData(ctx, *genState.LotteryData)
	}
	// start first lottery
	storedLottery := types.StoredLottery {
		Index: "0",
		Winner: "",
		WinIndex: 0,
		StartTxInd: 0,
	}
	k.SetStoredLottery(ctx, storedLottery)
	// update lotteryID
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if found {
		lotteryInfo.NextId++
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all lotteryInfo
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if found {
		genesis.LotteryInfo = &lotteryInfo
	}
	genesis.StoredLotteryList = k.GetAllStoredLottery(ctx)
	genesis.StoredBetList = k.GetAllStoredBet(ctx)
	// Get all betInfo
	betInfo, found := k.GetBetInfo(ctx)
	if found {
		genesis.BetInfo = &betInfo
	}
	// Get all lotteryData
	lotteryData, found := k.GetLotteryData(ctx)
	if found {
		genesis.LotteryData = &lotteryData
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
