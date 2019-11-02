package config

import (
	"errors"
	"log"
)

var (
	ErrUndefinedKey         = errors.New("undefined key")
	ErrRequiredKeysNotFound = errors.New("can not find required config keys")
)

type Config struct {
	values map[string]interface{}
}

func NewConfig() *Config {
	return &Config{
		values: make(map[string]interface{}),
	}
}

func (c *Config) Set(key string, value interface{}) {
	c.values[key] = value
}

func (c *Config) Unset(key string) error {
	if _, ok := c.values[key]; !ok {
		return ErrUndefinedKey
	}

	delete(c.values, key)

	return nil
}

func (c *Config) Get(key string) (value interface{}, ok bool) {
	value, ok = c.values[key]

	return
}

func (c *Config) GetString(key string) (string, bool) {
	v, ok := c.values[key]

	if !ok {
		return "", false
	}

	s, ok := v.(string)

	return s, ok
}

func (c *Config) RequireKeys(keys []string) error {
	hasErrors := false

	for _, k := range keys {
		if _, ok := c.values[k]; !ok {
			hasErrors = true

			log.Printf("Config with key %s is undefined", k)
		}
	}

	if hasErrors {
		return ErrRequiredKeysNotFound
	}

	return nil
}
