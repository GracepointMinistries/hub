package admin

import (
	"context"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/spf13/cobra"
)

var impersonateCmd = &cobra.Command{
	Use:   "impersonate",
	Short: "Impersonate the given user",
	Run: func(cmd *cobra.Command, args []string) {
		client := clientext.NewClient()
		payload, _, err := client.AdminApi.Impersonate(context.Background(), user)
		utils.CheckError(err)
		clientext.UpdateToken(payload.Token)
		clientext.WriteConfigFile()
	},
}

func init() {
	impersonateCmd.Flags().Int64VarP(&user, "user-id", "u", 0, "user to impersonate (required)")
	impersonateCmd.MarkFlagRequired("user-id")
}
