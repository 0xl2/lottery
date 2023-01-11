package keeper

import (
	"fmt"
	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bk types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CheckBalance(ctx sdk.Context, userAddr string) bool {
	// Get lotteryData
	lotteryData, found := k.GetLotteryData(ctx)
	if !found {
		panic("LotteryData not found")
	}

	userBalance := k.bankKeeper.SpendableCoins(ctx, sdk.AccAddress(userAddr)).AmountOf("demon");

	return userBalance.GTE(sdkmath.NewInt(int64(lotteryData.GetMinBet())))
}

func (k Keeper) CheckRange(ctx sdk.Context, betAmt uint64) bool {
	// Get lotteryData
	lotteryData, found := k.GetLotteryData(ctx)
	if !found {
		panic("LotteryData not found")
	}

	// check betAmt in min and max bet amount
	return betAmt >= lotteryData.GetMinBet() && betAmt <= lotteryData.GetMaxBet()
}

func (k Keeper) CheckBet(ctx sdk.Context, userAddr string) bool {
	// Get lotteryInfo
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if !found {
		panic("LotteryInfo not found")
	}

	// get current lottery from lotteryID
	lotteryID := lotteryInfo.GetNextId()
	currentLottery, found := k.GetStoredLottery(ctx, string(rune(lotteryID)))
	if !found {
		panic("Lottery not found")
	}

	// check user beted already
	starttxID := currentLottery.GetStartTxInd()
	lastTxID := starttxID + lotteryInfo.GetNextOrder()

	beted := false;
	for i := starttxID; i < lastTxID; i++ {
		selTx, found := k.GetStoredBet(ctx, string(rune(i)))
		if !found {
			panic("Tx not found")
		}

		if selTx.GetUserAddr() == userAddr {
			beted = true
			break
		}
	}

	return beted
}
