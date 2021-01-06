package cli

import (
	"github.com/InjectiveLabs/sdk-go/chain/erc20bridge/types"
	"github.com/InjectiveLabs/sdk-go/ethereum/rpc"
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
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBridgesRequest{}

			res, err := queryClient.Bridges(rpc.ContextWithHeight(clientCtx.Height), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
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
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryHubParamsRequest{}

			res, err := queryClient.HubParams(rpc.ContextWithHeight(clientCtx.Height), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
