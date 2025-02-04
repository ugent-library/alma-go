package main

import (
	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
)

var getBibsParams = alma.GetBibsParams{}

func init() {
	rootCmd.AddCommand(getBibsCmd)
	getBibsCmd.Flags().StringSliceVar(&getBibsParams.MmsID, "mms-id", nil, "")
	getBibsCmd.Flags().StringVar(&getBibsParams.IeID, "ie-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.HoldingsID, "holdings-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.RepresentationID, "representation-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.NzMmsID, "nz-mms-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.CzMmsID, "cz-mms-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.View, "view", "full", `"full" or "brief"`)
	getBibsCmd.Flags().StringSliceVar(&getBibsParams.Expand, "expand", nil, "")
	getBibsCmd.Flags().StringVar(&getBibsParams.OtherSystemID, "other-system-id", "", "")
	getBibsCmd.Flags().StringVar(&getBibsParams.LodUri, "lod-uri", "", "")
	getBibsCmd.MarkFlagsOneRequired("mms-id", "ie-id", "holdings-id", "representation-id", "nz-mms-id", "cz-mms-id", "other-system-id")
}

var getBibsCmd = &cobra.Command{
	Use:   "bibs",
	Short: "Get bibs",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetBibs(cmd.Context(), getBibsParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
