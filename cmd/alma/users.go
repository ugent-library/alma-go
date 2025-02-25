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
	Long: `Get users

# Retrieve the first batch of users
alma users 

# Retrieve the second batch of users
alma users --offset 10

# Search users based on last_name
alma users -q 'last_name~steenlant'

# Boolean search
alma users -q 'first_name~nicolas AND last_name~steenlant'

# Other search parameters: primary_id, first_name, last_name, middle_name, email, phone_number, job_category, identifiers, birth_date, user_group, campus_code, block_type, id_type, note_text, note_type, statistic_category, fines_fees_sum, general_info`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		almaClient := newAlmaClient()

		resBody, err := almaClient.RawGetUsers(cmd.Context(), getUsersParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
