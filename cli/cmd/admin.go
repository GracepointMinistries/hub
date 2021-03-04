package cmd

import (
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Operate on administrator resources",
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
