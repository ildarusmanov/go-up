package config

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

func (c *Config) Get(key string) (value interface{}, ok bool) {
	value, ok = c.values[key]

	return
}
