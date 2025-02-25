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
	Long: `Get library

# Retrieve information about the WE55 library in a pretty format
alma library WE55 --pretty`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetLibrary(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var getLibraryOpenHoursCmd = &cobra.Command{
	Use:   "open-hours [id]",
	Short: "Get library opening hours",
	Long: `Get library opening hours

# Retrieve the opening hours of the WE55 library in a pretty format
alma library open-hours WE55 --pretty`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetLibraryOpenHours(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
