package main

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
)

var deleteUserParams = alma.DeleteUserParams{}
var getUserFeesParams = alma.GetUserFeesParams{}

func init() {
	rootCmd.AddCommand(getUserCmd)

	getUserCmd.AddCommand(updateUserCmd)

	deleteUserCmd.Flags().StringVar(&deleteUserParams.UserIDType, "user-id-type", "", "")
	getUserCmd.AddCommand(deleteUserCmd)

	getUserFeesCmd.Flags().StringVar(&getUserFeesParams.UserIDType, "user-id-type", "", "")
	getUserFeesCmd.Flags().StringVar(&getUserFeesParams.Status, "status", "", "")
	getUserCmd.AddCommand(getUserFeesCmd)
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

var deleteUserCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		return almaClient.DeleteUser(cmd.Context(), args[0], deleteUserParams)
	},
}

var getUserFeesCmd = &cobra.Command{
	Use:   "fees [id]",
	Short: "Get user",
	Long: `Get user

# Retrieve user fees by user identifier
alma user fees 4685821335 > /tmp/fees.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetUserFees(cmd.Context(), args[0], alma.GetUserFeesParams{})
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
