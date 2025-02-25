package main

import (
	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
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
	Long: `Get printouts

# Retrieve the first 10 printouts
alma printouts

# Retrieve the second batch of 10 printouts
alma printouts --offset 10`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetPrintouts(cmd.Context(), getPrintoutsParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
