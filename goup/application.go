package goup

type Application interface {
	Restart() error
	Stop() error
}

type ServiceConfig interface{}

// Factory builds a service
// ctx contains value with key Application
// it should return a pointer to service instance and error
type ServiceFactory func(scfg ServiceConfig) (interface{}, error)

type ConfigManager interface {
	Set(key string, value interface{})
	Unset(string) error

	Get(key string) (interface{}, bool)
	GetString(key string) (string, bool)
	GetInt64(key string) (int64, bool)
	GetFloat64(key string) (float64, bool)
	GetBool(key string) (bool, bool)

	GetDefault(key string, defVal interface{}) interface{}
	GetDefaultString(key string, defVal string) string
	GetDefaultInt64(key string, defVal int64) int64
	GetDefaultFloat64(key string, defVal float64) float64
	GetDefaultBool(key string, defVal bool) bool

	RequireKeys([]string) error
}
