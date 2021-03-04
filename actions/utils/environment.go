package utils

import (
	"github.com/gobuffalo/envy"
)

// Environment wraps the environment that's set
type Environment string

const (
	developmentEnvironment Environment = "development"
	testEnvironment        Environment = "test"
	productionEnvironment  Environment = "production"
)

// GetEnvironment returns the current environment
func GetEnvironment() Environment {
	return Environment(envy.Get("GO_ENV", "development"))
}

// IsDeployed checks whether we're in a production deployed environment
func (e Environment) IsDeployed() bool {
	return e != developmentEnvironment && e != testEnvironment
}

// Load loads a dotenv file for the given environment
func (e Environment) Load() {
	envy.Load(".env." + string(e))
}
