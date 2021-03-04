package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Retrieve the profile of the scoped user",
	Run: func(cmd *cobra.Command, args []string) {
		payload, response, err := newClient().UserApi.Profile(context.Background())
		checkUnauthorized(response)
		checkError(err)
		dumpUsersWithZgroup(*payload.User)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
