package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/spf13/cobra"
)

func CmdShowLotteryInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-lottery-info",
		Short: "shows lotteryInfo",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLotteryInfoRequest{}

			res, err := queryClient.LotteryInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
