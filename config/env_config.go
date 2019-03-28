package config

import (
	"errors"
	"log"
	"os"
)

type EnvConfig struct{}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{}
}

func (c *EnvConfig) Set(key string, value interface{}) {
	os.Setenv(key, value.(string))
}

func (c *EnvConfig) Unset(key string) error {
	return os.Unsetenv(key)
}

func (c *EnvConfig) Get(key string) (value interface{}, ok bool) {
	return c.GetString(key)
}

func (c *EnvConfig) GetString(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (c *EnvConfig) RequireKeys(keys []string) error {
	hasErrors := false

	for _, k := range keys {
		if _, ok := os.LookupEnv(k); !ok {
			hasErrors = true

			log.Printf("Config with key %s is undefined", k)
		}
	}

	if hasErrors {
		return errors.New("Required keys does not exist in config")
	}

	return nil
}
