package goup

import (
	"context"
)

// Factory builds a service
// ctx contains value with key Application
// it should return a pointer to service instance and error
type ServiceFactory func(ctx context.Context) (interface{}, error)

type StopApplicationHandler func()

type ConfigManager interface {
	Set(key string, value interface{})
	Unset(string) error
	Get(key string) (interface{}, bool)
	GetString(key string) (string, bool)
	RequireKeys([]string) error
}
