package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getLibrariesCmd)
}

var getLibrariesCmd = &cobra.Command{
	Use:   "libraries",
	Short: "Get libraries",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetLibraries(cmd.Context())
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
