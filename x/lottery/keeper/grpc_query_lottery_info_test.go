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

func TestLotteryInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLotteryInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLotteryInfoRequest
		response *types.QueryGetLotteryInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotteryInfoRequest{},
			response: &types.QueryGetLotteryInfoResponse{LotteryInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LotteryInfo(wctx, tc.request)
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
