package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getLibraryCmd)
	getLibraryCmd.AddCommand(getLibraryOpenHoursCmd)
}

var getLibraryCmd = &cobra.Command{
	Use:   "library [id]",
	Short: "Get library",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetLibrary(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var getLibraryOpenHoursCmd = &cobra.Command{
	Use:   "open-hours [id]",
	Short: "Get library",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetLibraryOpenHours(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
