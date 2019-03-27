package goup

import (
	"context"

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
}

type DependenciesManager interface {
	Add(name string, service interface{})
	Get(name string) (interface{}, error)
}

type Application struct {
	ctx          context.Context
	Config       ConfigManager
	Dependencies DependenciesManager
}

// Creates new application
func NewApplication(ctx context.Context, d DependenciesManager, c ConfigManager) *Application {
	if d == nil {
		d = dependencies.NewDependencies()
	}

	if c == nil {
		c = config.NewConfig()
	}

	a := &Application{
		ctx:          ctx,
		Dependencies: d,
		Config:       c,
	}

	return a
}

// Add new service factory and create a service with it
func (a *Application) AddServiceFactory(srvName string, f ServiceFactory) error {
	ctx := context.WithValue(a.ctx, "Application", a)

	s, err := f(ctx)

	if err != nil {
		return err
	}

	a.Dependencies.Add(srvName, s)

	return nil
}

// Find service by name
func (a *Application) GetService(srvName string) (interface{}, error) {
	return a.Dependencies.Get(srvName)
}

// Set config variable
func (a *Application) SetConfig(key string, value interface{}) {
	a.Config.Set(key, value)
}

// Get config by key
func (a *Application) GetConfig(key string) (interface{}, bool) {
	return a.Config.Get(key)
}
