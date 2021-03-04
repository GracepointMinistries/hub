package cmd

import (
	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hub-cli",
	Short: "A Command Line Interface for the Hub API",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		clientext.InitializeClientConfig(cfgFile)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hub/cli.yaml)")
}
