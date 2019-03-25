package app

type Config struct {
	values map[string]string
}

func NewConfig() *Config {
	return &Config{
		values: make(map[string]string),
	}
}

func (c *Config) Set(key, value string) {
	c.values[key] = value
}

func (c *Config) Get(key string) (value string, ok bool) {
	value, ok = c.values[key]

	return
}
