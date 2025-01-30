package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugentlib/alma-go"
)

var getPrintoutsParams = alma.GetPrintoutsParams{}

func init() {
	rootCmd.AddCommand(getPrintoutsCmd)
	getPrintoutsCmd.Flags().StringVar(&getPrintoutsParams.Letter, "letter", "", "")
	getPrintoutsCmd.Flags().StringVar(&getPrintoutsParams.Status, "status", "", "")
	getPrintoutsCmd.Flags().StringVar(&getPrintoutsParams.PrinterID, "printer-id", "", "")
	getPrintoutsCmd.Flags().StringVar(&getPrintoutsParams.PrintedBy, "printed-by", "", "")
	getPrintoutsCmd.Flags().StringVar(&getPrintoutsParams.PrintoutID, "printout-id", "", "")
	getPrintoutsCmd.Flags().IntVar(&getPrintoutsParams.Limit, "limit", 10, "")
	getPrintoutsCmd.Flags().IntVar(&getPrintoutsParams.Offset, "offset", 0, "")
}

var getPrintoutsCmd = &cobra.Command{
	Use:   "printouts",
	Short: "Get printouts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetPrintouts(ctx, getPrintoutsParams)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}
