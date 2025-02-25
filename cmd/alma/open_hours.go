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
	Short: "Get open hours",
	Args:  cobra.NoArgs,
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
	Short: "Update open hours",
	Args:  cobra.NoArgs,
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
