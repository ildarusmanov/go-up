package main

import (
	"os"
	"strings"
)

type PackageDefinition struct {
	Name   string `yaml:"name,omitempty"`
	Import string `yaml:"import,omitempty"`
}

type Dependency struct {
	FactoryName       string             `yaml:"factory_name,omitempty"`
	Type              string             `yaml:"type,omitempty"`
	DependencyPackage *PackageDefinition `yaml:"dependency_package,omitempty"`
}

type ServiceFactory struct {
	FactoryName    string             `yaml:"factory_name,omitempty"`
	ServiceType    string             `yaml:"type_name,omitempty"`
	ServicePackage *PackageDefinition `yaml:"service_package,omitempty"`
	Dependencies   []*Dependency      `yaml:"dependencies,omitempty"`
}

func (s *ServiceFactory) FactoryFilename() string {
	return strings.ToLower(s.FactoryName) + "_factory.go"
}

type ServicesConfig struct {
	Pkgname  string            `yaml:"pkgname,omitempty"`
	Services []*ServiceFactory `yaml:"services,omitempty"`
}

func NewServicesConfig(pkgname string) *ServicesConfig {
	return &ServicesConfig{
		Pkgname:  pkgname,
		Services: []*ServiceFactory{},
	}
}

func (cfg *ServicesConfig) AddService(serviceFactory *ServiceFactory) {
	cfg.Services = append(cfg.Services, serviceFactory)
}

func (cfg *ServicesConfig) DropService(wdir, factoryName string) {
	newServices := []*ServiceFactory{}

	for _, v := range cfg.Services {
		if v.FactoryName == factoryName {
			os.Remove(wdir + "/app/" + v.FactoryFilename())

			continue
		}

		newServices = append(newServices, v)
	}

	cfg.Services = newServices
}
