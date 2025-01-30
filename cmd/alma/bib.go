package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugentlib/alma-go"
)

var getBibParams = alma.GetBibParams{}

var getItemsParams = alma.GetHoldingItemsParams{}

func init() {
	rootCmd.AddCommand(getBibCmd)
	getBibCmd.Flags().StringVar(&getBibParams.View, "view", "full", `"full" or "brief"`)
	getBibCmd.Flags().StringSliceVar(&getBibParams.Expand, "expand", nil, "")

	getBibCmd.AddCommand(getHoldingsCmd)

	getBibCmd.AddCommand(getHoldingCmd)

	getHoldingCmd.AddCommand(getHoldingItemsCmd)
	getHoldingItemsCmd.Flags().IntVar(&getItemsParams.Limit, "limit", 10, "")
	getHoldingItemsCmd.Flags().IntVar(&getItemsParams.Offset, "offset", 0, "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.Expand, "expand", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.UserID, "user-id", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CurrentLibrary, "current-library", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CurrentLocation, "current-location", "", "")
	getHoldingItemsCmd.Flags().StringVarP(&getItemsParams.Query, "query", "q", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.OrderBy, "order-by", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.Direction, "direction", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CreateDateFrom, "create-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CreateDateTo, "create-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ModifyDateFrom, "modify-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ModifyDateTo, "modify-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ReceiveDateFrom, "receive-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ReceiveDateTo, "receive-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ExpectedReceiveDateFrom, "expected-receive-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ExpectedReceiveDateTo, "expected-receive-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.View, "view", "brief", `"full" or "brief"`)
}

var getBibCmd = &cobra.Command{
	Use:   "bib [mms-id]",
	Short: "Get bib",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetBib(ctx, args[0], getBibParams)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}

var getHoldingsCmd = &cobra.Command{
	Use:   "holdings [mms-id]",
	Short: "Get bib holdings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetHoldings(ctx, args[0])
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}

var getHoldingCmd = &cobra.Command{
	Use:   "holding [mms-id] [holding-id]",
	Short: "Get bib holding",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetHolding(ctx, args[0], args[1])
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}

var getHoldingItemsCmd = &cobra.Command{
	Use:   "items [mms-id] [holding-id]",
	Short: "Get bib holding items",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetHoldingItems(ctx, args[0], args[1], getItemsParams)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}
