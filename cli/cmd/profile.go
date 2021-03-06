package cmd

import (
	"context"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/print"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Retrieve the profile of the scoped user",
	Run: func(cmd *cobra.Command, args []string) {
		payload, response, err := clientext.NewClient().UserApi.Profile(context.Background())
		utils.CheckUnauthorized(response)
		utils.CheckError(err)
		print.DumpUsers(*payload.User)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
