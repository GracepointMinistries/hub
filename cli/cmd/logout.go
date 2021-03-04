package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "log-out",
	Short: "Stop impersonating a user",
	Run: func(cmd *cobra.Command, args []string) {
		payload, response, err := newClient().UserApi.Logout(context.Background())
		checkUnauthorized(response)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fileConfig.Token = payload.Token
		writeConfigFile()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
