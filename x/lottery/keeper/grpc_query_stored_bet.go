package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/naruto0913/lottery/x/lottery/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredBetAll(c context.Context, req *types.QueryAllStoredBetRequest) (*types.QueryAllStoredBetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedBets []types.StoredBet
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedBetStore := prefix.NewStore(store, types.KeyPrefix(types.StoredBetKeyPrefix))

	pageRes, err := query.Paginate(storedBetStore, req.Pagination, func(key []byte, value []byte) error {
		var storedBet types.StoredBet
		if err := k.cdc.Unmarshal(value, &storedBet); err != nil {
			return err
		}

		storedBets = append(storedBets, storedBet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredBetResponse{StoredBet: storedBets, Pagination: pageRes}, nil
}

func (k Keeper) StoredBet(c context.Context, req *types.QueryGetStoredBetRequest) (*types.QueryGetStoredBetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredBet(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredBetResponse{StoredBet: val}, nil
}
