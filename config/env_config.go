package config

import (
	"errors"
	"log"
	"os"
	"strconv"
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

func (c *EnvConfig) GetInt64(key string) (r int64, ok bool) {
	var (
		v   string
		err error
	)

	v, ok = c.GetString(key)

	if !ok {
		return
	}

	r, err = strconv.ParseInt(v, 10, 64)

	if err != nil {
		return
	}

	return
}

func (c *EnvConfig) GetFloat64(key string) (r float64, ok bool) {
	var (
		v   string
		err error
	)

	v, ok = c.GetString(key)

	if !ok {
		return
	}

	r, err = strconv.ParseFloat(v, 64)

	if err != nil {
		return
	}

	return
}

func (c *EnvConfig) GetBool(key string) (r bool, ok bool) {
	var (
		v   string
		err error
	)

	v, ok = c.GetString(key)

	if !ok {
		return
	}

	r, err = strconv.ParseBool(v)

	if err != nil {
		return
	}

	return
}

func (c *EnvConfig) GetDefault(key string, defVal interface{}) (value interface{}) {
	if v, ok := c.Get(key); ok {
		return v
	}

	return defVal
}

func (c *EnvConfig) GetDefaultString(key string, defVal string) string {
	if v, ok := c.GetString(key); ok {
		return v
	}

	return defVal
}

func (c *EnvConfig) GetDefaultInt64(key string, defVal int64) int64 {
	if v, ok := c.GetInt64(key); ok {
		return v
	}

	return defVal
}

func (c *EnvConfig) GetDefaultFloat64(key string, defVal float64) float64 {
	if v, ok := c.GetFloat64(key); ok {
		return v
	}

	return defVal
}

func (c *EnvConfig) GetDefaultBool(key string, defVal bool) bool {
	if v, ok := c.GetBool(key); ok {
		return v
	}

	return defVal
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
