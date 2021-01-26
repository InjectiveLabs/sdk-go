package cli

import (
	"context"

	"github.com/InjectiveLabs/sdk-go/chain/erc20bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all modules/bank CLi query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc20bridge module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetBridgesCmd(),
		GetHubParamsCmd(),
	)
	return cmd
}

// GetBridgesCmd queries a bridges registered
func GetBridgesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridges",
		Short: "Gets bridges registered",
		Long:  "Gets bridges registered.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBridgesRequest{}

			res, err := queryClient.Bridges(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetHubParamsCmd queries hub info
func GetHubParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hubparams",
		Short: "Gets hub info.",
		Long:  "Gets hub info.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryHubParamsRequest{}

			res, err := queryClient.HubParams(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
