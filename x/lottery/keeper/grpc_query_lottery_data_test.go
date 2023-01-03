package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/naruto0913/lottery/testutil/keeper"
	"github.com/naruto0913/lottery/testutil/nullify"
	"github.com/naruto0913/lottery/x/lottery/types"
)

func TestLotteryDataQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLotteryData(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLotteryDataRequest
		response *types.QueryGetLotteryDataResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotteryDataRequest{},
			response: &types.QueryGetLotteryDataResponse{LotteryData: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LotteryData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
