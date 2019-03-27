package dependencies

import (
	"errors"
	"sync"
)

var (
	serviceNotFound = errors.New("service not found")
)

type Dependencies struct {
	items map[string]interface{}
	mux   sync.Mutex
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		items: make(map[string]interface{}),
	}
}

func (d *Dependencies) Add(name string, service interface{}) {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.items[name] = service
}

func (d *Dependencies) Get(name string) (interface{}, error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	s, ok := d.items[name]

	if !ok {
		return nil, serviceNotFound
	}

	return s, nil
}
