package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
)

var (
	getOpenHoursParams    = alma.GetOpenHoursParams{}
	updateOpenHoursParams = alma.UpdateOpenHoursParams{}
)

func init() {
	rootCmd.AddCommand(getOpenHoursCmd)
	getOpenHoursCmd.Flags().StringVar(&getOpenHoursParams.Scope, "scope", "", "")
	getOpenHoursCmd.AddCommand(updateOpenHoursCmd)
	updateOpenHoursCmd.Flags().StringVar(&updateOpenHoursParams.Scope, "scope", "", "")
	updateOpenHoursCmd.MarkFlagRequired("scope")
}

var getOpenHoursCmd = &cobra.Command{
	Use:   "open-hours",
	Short: "Get openings hours",
	Long: `Get openings hours

# Retrieve the opening hours for the WE55 library
alma open-hours --scope WE55 > /tmp/WE.hours.txt`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetOpenHours(cmd.Context(), getOpenHoursParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var updateOpenHoursCmd = &cobra.Command{
	Use:   "update",
	Short: "Update openings hours",
	Long: `Update opening hours

# Update the opening hours for the WE55 library
alma open-hours update --scope WE55 < /tmp/WE.hours.txt`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		reqBody, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		resBody, err := almaClient.RawUpdateOpenHours(cmd.Context(), updateOpenHoursParams, reqBody)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
