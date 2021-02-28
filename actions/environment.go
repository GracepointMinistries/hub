package actions

import (
	"github.com/gobuffalo/envy"
)

type environment string

const (
	developmentEnvironment environment = "development"
	testEnvironment        environment = "test"
	productionEnvironment  environment = "production"
)

func getEnvironment() environment {
	return environment(envy.Get("GO_ENV", "development"))
}

func (e environment) IsDeployed() bool {
	return e != developmentEnvironment && e != testEnvironment
}

func (e environment) Load() {
	envy.Load(".env." + string(e))
}
