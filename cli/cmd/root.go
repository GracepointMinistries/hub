package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type config struct {
	Token string `yaml:"token"`
	Host  string `yaml:"host"`
}

var (
	cfgFile    string
	fileConfig config
	rootCmd    = &cobra.Command{
		Use:   "hub-cli",
		Short: "A Command Line Interface for the Hub API",
	}
)

func writeConfigFile() {
	data, err := yaml.Marshal(&fileConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	err = os.MkdirAll(path.Dir(cfgFile), 0700)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(cfgFile, data, 0600)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hub/cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hub/cli.yaml" (without extension).
		basePath := path.Join(home, ".hub")
		viper.AddConfigPath(basePath)
		viper.SetConfigName("cli")
		cfgFile = path.Join(basePath, "cli.yaml")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

	if err := viper.Unmarshal(&fileConfig); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
