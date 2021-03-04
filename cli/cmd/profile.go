package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Retrieve the profile of the scoped user",
	Run: func(cmd *cobra.Command, args []string) {
		payload, response, err := newClient().UserApi.Profile(context.Background())
		checkUnauthorized(response)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		dumpUser(payload.User, payload.Zgroup, true)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
