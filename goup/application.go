package goup

import (
	"context"
	"log"

	"github.com/ildarusmanov/go-up/config"
	"github.com/ildarusmanov/go-up/dependencies"
)

// Factory builds a service
// ctx contains value with key Application
// it should return a pointer to service instance and error
type ServiceFactory func(ctx context.Context) (interface{}, error)

// Service is just an empty interface
type Service interface{}

type ConfigManager interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	GetString(key string) (string, bool)
}

type DependenciesManager interface {
	Add(name string, service interface{})
	Get(name string) (interface{}, error)
}

type Application struct {
	ctx  context.Context
	cfg  ConfigManager
	deps DependenciesManager
}

// Creates new application
func NewApplication() *Application {
	return &Application{
		ctx:  context.Background(),
		deps: dependencies.NewDependencies(),
		cfg:  config.NewConfig(),
	}
}

// Set dependency manager
func (a *Application) WithDependencies(d DependenciesManager) *Application {
	if d == nil {
		log.Fatalf("WithDependencies(%v) failed", d)
	}

	a.deps = d

	return a
}

// Get dependency manager
func (a *Application) Dependencies() DependenciesManager {
	return a.deps
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

// Add new service factory and create a service with it
func (a *Application) AddServiceFactory(srvName string, f ServiceFactory) error {
	ctx := context.WithValue(a.Context(), "Application", a)

	s, err := f(ctx)

	if err != nil {
		return err
	}

	a.Dependencies().Add(srvName, s)

	return nil
}

// Find service by name
func (a *Application) GetService(srvName string) (interface{}, error) {
	return a.Dependencies().Get(srvName)
}

// Set config variable
func (a *Application) SetConfig(key string, value interface{}) {
	a.Config().Set(key, value)
}

// Get config by key
func (a *Application) GetConfig(key string) (interface{}, bool) {
	return a.Config().Get(key)
}

// Get config string by key
func (a *Application) GetConfigString(key string) (string, bool) {
	return a.Config().GetString(key)
}
