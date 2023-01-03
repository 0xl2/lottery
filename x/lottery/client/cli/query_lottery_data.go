package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/naruto0913/lottery/x/lottery/types"
	"github.com/spf13/cobra"
)

func CmdShowLotteryData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-lottery-data",
		Short: "shows lotteryData",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLotteryDataRequest{}

			res, err := queryClient.LotteryData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
