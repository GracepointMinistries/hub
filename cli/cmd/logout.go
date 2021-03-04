package cmd

import (
	"context"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "log-out",
	Short: "Stop impersonating a user",
	Run: func(cmd *cobra.Command, args []string) {
		payload, response, err := clientext.NewClient().UserApi.Logout(context.Background())
		utils.CheckUnauthorized(response)
		utils.CheckError(err)
		clientext.UpdateToken(payload.Token)
		clientext.WriteConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
