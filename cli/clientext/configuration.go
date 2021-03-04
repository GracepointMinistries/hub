package clientext

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/GracepointMinistries/hub/cli/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type config struct {
	Token string `yaml:"token"`
	Host  string `yaml:"host"`
}

var (
	// holds the unmarshaled configuration
	fileConfig    config
	globalCfgFile string
)

// UpdateToken updates the globally stored token
func UpdateToken(token string) {
	fileConfig.Token = token
}

// UpdateHost updates the globally stored host
func UpdateHost(host string) {
	fileConfig.Host = host
}

// WriteConfigFile flushes any config file changes to disk
func WriteConfigFile() {
	data, err := yaml.Marshal(&fileConfig)
	utils.CheckError(err)
	utils.CheckError(os.MkdirAll(path.Dir(globalCfgFile), 0700))
	utils.CheckError(ioutil.WriteFile(globalCfgFile, data, 0600))
}

// InitializeClientConfig initializes the global client configuration
// based off of a config file, it needs to be called prior to
// initializing a client or flushing state to disk
func InitializeClientConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		utils.CheckError(err)

		// Search config in home directory with name ".hub/cli.yaml" (without extension).
		basePath := path.Join(home, ".hub")
		viper.AddConfigPath(basePath)
		viper.SetConfigName("cli")
		cfgFile = path.Join(basePath, "cli.yaml")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
	utils.CheckError(viper.Unmarshal(&fileConfig))
	globalCfgFile = cfgFile
}
