package config

import "sync"

type Config struct {
	AppConfig *appConfig
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		cfg = Config{
			AppConfig: newappConfig(),
		}
	})
	return &cfg
}
