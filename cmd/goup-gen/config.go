package main

import (
	"os"
	"strings"
)

type PackageDefinition struct {
	Name   string `yaml:"name,omitempty"`
	Import string `yaml:"import,omitempty"`
}

func (p *PackageDefinition) GetDefinition() string {
	if p.Name == "" {
		return `"` + p.Import + `"`
	}

	return p.Name + ` "` + p.Import + `"`
}

type Dependency struct {
	FactoryName       string             `yaml:"factory_name,omitempty"`
	Type              string             `yaml:"type,omitempty"`
	DependencyPackage *PackageDefinition `yaml:"dependency_package,omitempty"`
}

type Factory struct {
	FactoryName        string               `yaml:"factory_name,omitempty"`
	ServiceType        string               `yaml:"type_name,omitempty"`
	ServicePackage     *PackageDefinition   `yaml:"service_package,omitempty"`
	ServiceConstructor *Constructor         `yaml:"service_constructor,omitempty"`
	Dependencies       []*Dependency        `yaml:"dependencies,omitempty"`
	Imports            []*PackageDefinition `yaml:"imports,omitempty"`
	FactoryConfig      *FactoryConfig       `yaml:"factory_config,omitempty"`
}

type FactoryConfig struct {
	Fields []*FactoryConfigField `yaml:"fields,omitempty"`
}

type FactoryConfigField struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
}

type Constructor struct {
	Signature string `yaml:"signature,omitempty"`
}

func (s *Factory) GetFactoryConfigFields() []*FactoryConfigField {
	if s.FactoryConfig == nil {
		return []*FactoryConfigField{}
	}

	return s.FactoryConfig.Fields
}

func (s *Factory) FactoryFilename() string {
	return strings.ToLower(s.FactoryName) + "_factory.go"
}

type ServicesConfig struct {
	Pkgname   string     `yaml:"pkgname,omitempty"`
	Factories []*Factory `yaml:"factories,omitempty"`
}

func NewServicesConfig(pkgname string) *ServicesConfig {
	return &ServicesConfig{
		Pkgname:   pkgname,
		Factories: []*Factory{},
	}
}

func (cfg *ServicesConfig) AddFactory(factory *Factory) {
	cfg.Factories = append(cfg.Factories, factory)
}

func (cfg *ServicesConfig) DropFactory(wdir, factoryName string) {
	newFactories := []*Factory{}

	for _, v := range cfg.Factories {
		if v.FactoryName == factoryName {
			os.Remove(wdir + "/app/" + v.FactoryFilename())

			continue
		}

		newFactories = append(newFactories, v)
	}

	cfg.Factories = newFactories
}
