package config

import "github.com/kelseyhightower/envconfig"

type appConfig struct {
	Port        string `envconfig:"PORT" default:"8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"DEV"`
}

func newappConfig() *appConfig {
	var appCfg appConfig
	envconfig.MustProcess("", &appCfg)
	return &appCfg
}
