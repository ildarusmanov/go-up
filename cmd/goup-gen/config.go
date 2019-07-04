package main

type ServiceFactory struct {
	ServiceName string `yaml:"service_name,omitempty"`
	FactoryName string `yaml:"factory_name,omitempty"`
	MethodName  string `yaml:"method_name,omitempty"`
	ServiceType string `yaml:"type_name,omitempty"`
	Filename    string `yaml:"filename,omitempty"`
}

type ServicesConfig struct {
	Services []*ServiceFactory `yaml:"services,omitempty"`
}
