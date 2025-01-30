package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugentlib/alma-go"
)

var getRequestedResourcesParams = alma.GetRequestedResourcesParams{}

func init() {
	rootCmd.AddCommand(getRequestedResourcesCmd)
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.CircDesk, "circ-desk", "", "e.g. DEFAULT_CIRC_DESK")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.Library, "library", "", "e.g. MAIN")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.Location, "location", "", "")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.OrderBy, "order-by", "", "")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.Direction, "direction", "", "")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.PickupInst, "pickup-inst", "", "")
	getRequestedResourcesCmd.Flags().StringVar(&getRequestedResourcesParams.Reported, "reported", "", "")
	getRequestedResourcesCmd.Flags().IntVar(&getRequestedResourcesParams.Limit, "limit", 10, "")
	getRequestedResourcesCmd.Flags().IntVar(&getRequestedResourcesParams.Offset, "offset", 0, "")
	getRequestedResourcesCmd.MarkFlagRequired("circ-desk")
	getRequestedResourcesCmd.MarkFlagRequired("library")
}

var getRequestedResourcesCmd = &cobra.Command{
	Use:   "requested-resources",
	Short: "Get requested resources",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetRequestedResources(ctx, getRequestedResourcesParams)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}
