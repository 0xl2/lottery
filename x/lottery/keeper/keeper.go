package keeper

import (
	"fmt"
	"context"
	"strconv"
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

func (k Keeper) getWinner(ctx sdk.Context, txCnt int64) (int64, error) {
	return 1, nil
}

// function to check if beter has sufficient balance
func (k Keeper) CheckBalance(ctx sdk.Context, userAddr string) (bool, error) {
	// Get lotteryData
	lotteryData, found := k.GetLotteryData(ctx)
	if !found {
		return false, types.ErrLotteryData
	}

	userBalance := k.bankKeeper.SpendableCoins(ctx, sdk.AccAddress(userAddr)).AmountOf(types.LottDenom);

	return userBalance.GTE(sdkmath.NewInt(int64(lotteryData.GetMinBet()))), nil
}

// function to check betAmount is in valid range
func (k Keeper) CheckRange(ctx sdk.Context, betAmt int64) (bool, error) {
	// Get lotteryData
	lotteryData, found := k.GetLotteryData(ctx)
	if !found {
		return false, types.ErrLotteryData
	}

	// check betAmt in min and max bet amount
	return betAmt >= lotteryData.GetMinBet() && betAmt <= lotteryData.GetMaxBet(), nil
}

// function to check if beter bettted or not
func (k Keeper) CheckBet(ctx sdk.Context, userAddr string) (bool, error) {
	// Get lotteryInfo
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if !found {
		return false, types.ErrLotteryInfo
	}

	// get current lottery from lotteryID
	lotteryID := lotteryInfo.GetNextId()
	newIndex := strconv.FormatInt(lotteryID, 10)
	currentLottery, found := k.GetStoredLottery(ctx, newIndex)
	if !found {
		return false, types.ErrLottery
	}

	// check user beted already
	starttxID := currentLottery.GetStartTxInd()
	lastTxID := starttxID + lotteryInfo.GetNextOrder()

	beted := false;
	for i := starttxID; i < lastTxID; i++ {
		iStr := strconv.FormatInt(i, 10)
		selTx, found := k.GetStoredBet(ctx, iStr)
		if !found {
			return false, types.ErrBetTx
		}

		if selTx.GetUserAddr() == userAddr {
			beted = true
			break
		}
	}

	return beted, nil
}

// function to save the bet tx and update lottery, bet
func (k Keeper) SaveBet(ctx sdk.Context, userAddr string, betAmt int64) (*types.MsgDoBetResponse, error) {
	// save bet info
	betInfo, found := k.GetBetInfo(ctx)
	if !found {
		return nil, types.ErrBetInfo
	}

	betID := betInfo.GetBetId()
	newIndex := strconv.FormatInt(betID, 10)
	storedBet := types.StoredBet{
		Index: newIndex,
		UserAddr: userAddr,
		BetAmount: betAmt,
	}
	k.SetStoredBet(ctx, storedBet)

	// transfertoken
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, sdk.AccAddress(userAddr), types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.LottDenom, sdkmath.NewInt(betAmt)))); err != nil {
		return &types.MsgDoBetResponse{
			BetOrder: newIndex,
		}, err
	}

	// get lottery info
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if !found {
		return nil, types.ErrLotteryInfo
	}

	ctx.EventManager().EmitTypedEvent(&types.DoBet{
		LotteryID: lotteryInfo.GetNextId(),
		BetID: betID,
		UserAddr: userAddr,
		BetAmount: betAmt,
	})

	// get current lottery
	newIndex = strconv.FormatInt(lotteryInfo.GetNextId(), 10)
	currentLottery, found := k.GetStoredLottery(ctx, newIndex)
	if !found {
		return nil, types.ErrLottery
	}
	
	// update current lottery totalAmt
	currentLottery.TotalAmt += betAmt

	// update lottery info
	lotteryInfo.NextOrder++

	// update betID
	betInfo.BetId++

	return &types.MsgDoBetResponse{
		BetOrder: newIndex,
	}, nil
}

func (k Keeper) SetWinner(goCtx context.Context) (error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get lotteryinfo
	lotteryInfo, found := k.GetLotteryInfo(ctx)
	if !found {
		return types.ErrLotteryInfo
	}

	lotteryID := lotteryInfo.GetNextId()
	// if lottery started
	if lotteryID != 0 {
		newIndex := strconv.FormatInt(lotteryID - 1, 10)	
		currentLottery, found := k.GetStoredLottery(ctx, newIndex)
		if !found {
			return types.ErrLottery
		}

		// get tx count
		txCnt := lotteryInfo.GetNextOrder()

		// complete lottery if tx count is over 10
		if txCnt >= 10 {
			starttxID := currentLottery.GetStartTxInd()

			// get winner and set winner in lottery
			winIndex, err := k.getWinner(ctx, txCnt)
			if err != nil {
				return err
			}

			iStr := strconv.FormatInt(starttxID + winIndex, 10)
			selTx, found := k.GetStoredBet(ctx, iStr)
			if !found {
				return types.ErrBetTx
			}

			// update current lottery info
			currentLottery.WinIndex = winIndex
			currentLottery.Winner = selTx.GetUserAddr()

			// Get lotteryData
			lotteryData, found := k.GetLotteryData(ctx)
			if !found {
				return types.ErrLotteryData
			}

			// send token to winner
			totalAmt := currentLottery.GetTotalAmt() - lotteryData.GetFee()
			if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(currentLottery.Winner), sdk.NewCoins(sdk.NewCoin(types.LottDenom, sdkmath.NewInt(totalAmt)))); err != nil {
				return err
			}

			// get bet info
			betInfo, found := k.GetBetInfo(ctx)
			if !found {
				return types.ErrBetInfo
			}
			// add new lottery
			newIndex := strconv.FormatInt(lotteryID, 10)
			storedLottery := types.StoredLottery {
				Index: newIndex,
				Winner: "",
				WinIndex: 0,
				StartTxInd: betInfo.GetBetId(),
			}
			k.SetStoredLottery(ctx, storedLottery)

			// update lottery info
			lotteryInfo.NextId++
			lotteryInfo.NextOrder = 0
		}
	}

	return nil
}
