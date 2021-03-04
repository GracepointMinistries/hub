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

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, criticalf("Error: %v", err))
		os.Exit(1)
	}
}

func writeConfigFile() {
	data, err := yaml.Marshal(&fileConfig)
	checkError(err)
	checkError(os.MkdirAll(path.Dir(cfgFile), 0700))
	checkError(ioutil.WriteFile(cfgFile, data, 0600))
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
		checkError(err)

		// Search config in home directory with name ".hub/cli.yaml" (without extension).
		basePath := path.Join(home, ".hub")
		viper.AddConfigPath(basePath)
		viper.SetConfigName("cli")
		cfgFile = path.Join(basePath, "cli.yaml")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
	checkError(viper.Unmarshal(&fileConfig))
}
