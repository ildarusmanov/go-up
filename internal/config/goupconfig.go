package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

const (
	GoupConfigFileName = ".goup.yml"
)

type GoupConfig struct {
	Pkgname string `yaml:"pkgname,omitempty"`
	Version string `yaml:"pkgname,omitempty"`
}

func NewGoupConfig(pkgname string) *GoupConfig {
	return &GoupConfig{
		Pkgname: pkgname,
	}
}

func LoadGoupConfig(pdir string) (*GoupConfig, error) {
	yamlFile := pdir + "/" + GoupConfigFileName

	f, err := os.Open(yamlFile)

	if err != nil {
		log.Printf("Can not open servces config file %s: %s\n", yamlFile, err)
		return nil, err
	}

	d := yaml.NewDecoder(f)

	cfg := &GoupConfig{}

	if err := d.Decode(cfg); err != nil {
		log.Printf("Can not parse services config %s: %s\n", yamlFile, err)
		return nil, err
	}

	return cfg, nil
}
