package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersCmd)
	usersCmd.AddCommand(getUserCmd)
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "",
}

var getUserCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// d, err := almaClient.GetUser(ctx, args[0])
		// if err != nil {
		// 	return err
		// }
		// b, err := json.Marshal(d)
		// if err != nil {
		// 	return err
		// }

		b, err := almaClient.GetRawUser(ctx, args[0])
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(b)

		return err
	},
}
