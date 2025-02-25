package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getUserCmd)
	getUserCmd.AddCommand(updateUserCmd)
}

var getUserCmd = &cobra.Command{
	Use:   "user [id]",
	Short: "Get user",
	Long: `Get user

# Retrieve a user by primary id
alma user 4685821335 > /tmp/user.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetUser(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var updateUserCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update user",
	Long: `Update user

# Update a user record with primary id 4685821335
alma user update 4685821335 < /tmp/user.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		reqBody, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		resBody, err := almaClient.RawUpdateUser(cmd.Context(), args[0], reqBody)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
