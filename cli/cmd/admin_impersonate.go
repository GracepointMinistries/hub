package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	user int64
)

var impersonateCmd = &cobra.Command{
	Use:   "impersonate",
	Short: "Impersonate the given user",
	Run: func(cmd *cobra.Command, args []string) {
		client := newClient()
		payload, _, err := client.AdminApi.Impersonate(context.Background(), user)
		checkError(err)
		fileConfig.Token = payload.Token
		writeConfigFile()
	},
}

func init() {
	impersonateCmd.Flags().Int64VarP(&user, "user-id", "u", 0, "user to impersonate (required)")
	impersonateCmd.MarkFlagRequired("user-id")

	adminCmd.AddCommand(impersonateCmd)
}
