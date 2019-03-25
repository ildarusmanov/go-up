package app

import "context"

// Factory builds a service
// ctx contains value with key Application
// it should return a pointer to service instance and error
type ServiceFactory func(ctx context.Context) (Service, error)

// Service is just an empty interface
type Service interface{}

type Application struct {
	ctx          context.Context
	Config       *Config
	Dependencies *Dependencies
}

// Creates new application
func NewApplication(ctx context.Context, d *Dependencies, c *Config) *Application {
	if d == nil {
		d = NewDependencies()
	}

	if c == nil {
		c = NewConfig()
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
func (a *Application) GetService(srvName string) (Service, error) {
	return a.Dependencies.Get(srvName)
}

// Set config variable
func (a *Application) SetConfig(key, value string) {
	a.Config.Set(key, value)
}

// Get config by key
func (a *Application) GetConfig(key string) (string, bool) {
	return a.Config.Get(key)
}
