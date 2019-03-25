package app

import "context"

type Application struct {
	ctx          context.Context
	Config       *Config
	Dependencies *Dependencies
}

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

func (a *Application) AddServiceFactory(srvName string, f ServiceFactory) error {
	ctx := context.WithValue(a.ctx, "Application", a)

	s, err := f(ctx)

	if err != nil {
		return err
	}

	a.Dependencies.Add(srvName, s)

	return nil
}

func (a *Application) GetService(srvName string) (Service, error) {
	return a.Dependencies.Get(srvName)
}

func (a *Application) SetConfig(key, value string) {
	a.Config.Set(key, value)
}

func (a *Application) GetConfig(key string) (string, bool) {
	return a.Config.Get(key)
}
