package cmd

import (
	"github.com/GracepointMinistries/hub/cli/cmd/admin"
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Operate on administrator resources",
}

func init() {
	admin.Register(adminCmd)
	rootCmd.AddCommand(adminCmd)
}
