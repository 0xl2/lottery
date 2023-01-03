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
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
