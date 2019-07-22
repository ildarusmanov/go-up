package main

type PackageDefinition struct {
	Name   string `yaml:"name,omitempty"`
	Import string `yaml:"import,omitempty"`
}

type ServiceFactory struct {
	ServiceName     string             `yaml:"service_name,omitempty"`
	FactoryName     string             `yaml:"factory_name,omitempty"`
	MethodName      string             `yaml:"method_name,omitempty"`
	ServiceType     string             `yaml:"type_name,omitempty"`
	ServicePackage  *PackageDefinition `yaml:"service_package,omitempty"`
	FactoryFilename string             `yaml:"factory_filename,omitempty"`
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
