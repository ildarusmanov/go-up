package goup

import (
	"context"
	"log"

	"github.com/ildarusmanov/go-up/config"
)

// Factory builds a service
// ctx contains value with key Application
// it should return a pointer to service instance and error
type ServiceFactory func(ctx context.Context) (interface{}, error)

// Service is just an empty interface
type Service interface{}

type ConfigManager interface {
	Set(key string, value interface{})
	Unset(string) error
	Get(key string) (interface{}, bool)
	GetString(key string) (string, bool)
	RequireKeys([]string) error
}

type Application struct {
	ctx context.Context
	cfg ConfigManager
}

// Creates new application
func NewApplication() *Application {
	return &Application{
		ctx: context.Background(),
		cfg: config.NewConfig(),
	}
}

// Set config manager
func (a *Application) WithConfig(c ConfigManager) *Application {
	if c == nil {
		log.Fatalf("WithConfig(%v) failed", c)
	}

	a.cfg = c

	return a
}

// Get config manager
func (a *Application) Config() ConfigManager {
	return a.cfg
}

// Set context
func (a *Application) WithContext(ctx context.Context) *Application {
	a.ctx = ctx

	return a
}

// Get current context
func (a *Application) Context() context.Context {
	return a.ctx
}

// Check config with keys
func (a *Application) RequireConfig(keys []string) error {
	return a.Config().RequireKeys(keys)
}

// Set config variable
func (a *Application) SetConfig(key string, value interface{}) {
	a.Config().Set(key, value)
}

// Remove config by key
func (a *Application) UnsetConfig(key string) error {
	return a.Config().Unset(key)
}

// Get config by key
func (a *Application) GetConfig(key string) (interface{}, bool) {
	return a.Config().Get(key)
}

// Get config string by key
func (a *Application) GetConfigString(key string) (string, bool) {
	return a.Config().GetString(key)
}
