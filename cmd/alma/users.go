package main

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(getUserCmd)
	userCmd.AddCommand(updateUserCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "",
}

var getUserCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		resBody, err := almaClient.RawGetUser(ctx, args[0])
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}

var updateUserCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		reqBody, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		resBody, err := almaClient.RawUpdateUser(ctx, args[0], reqBody)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(resBody)

		return err
	},
}
