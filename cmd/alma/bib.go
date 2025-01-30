package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugentlib/alma-go"
)

var getBibParams = alma.GetBibParams{}

func init() {
	rootCmd.AddCommand(getBibCmd)
	getBibCmd.AddCommand(getHoldingsCmd)
	getBibCmd.Flags().StringVar(&getBibParams.View, "view", "full", `"full" or "brief"`)
	getBibCmd.Flags().StringSliceVar(&getBibParams.Expand, "expand", nil, "")
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
	Use:   "holdings",
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
