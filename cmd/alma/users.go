package main

import (
	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
)

var getUsersParams = alma.GetUsersParams{}

func init() {
	rootCmd.AddCommand(getUsersCmd)
	getUsersCmd.Flags().StringVar(&getUsersParams.Expand, "expand", "", "")
	getUsersCmd.Flags().IntVar(&getUsersParams.Limit, "limit", 10, "")
	getUsersCmd.Flags().IntVar(&getUsersParams.Offset, "offset", 0, "")
	getUsersCmd.Flags().StringVar(&getUsersParams.OrderBy, "order-by", "", "")
	getUsersCmd.Flags().StringVarP(&getUsersParams.Query, "query", "q", "", "")
	getUsersCmd.Flags().StringVar(&getUsersParams.SourceInstitutionCode, "source-institution-code", "", "")
	getUsersCmd.Flags().StringVar(&getUsersParams.SourceUserID, "source-user-id", "", "")
}

var getUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get users",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetUsers(cmd.Context(), getUsersParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
