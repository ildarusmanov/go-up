package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func parseConfigYaml(servicesYamlFile string) (*ServicesConfig, error) {
	f, err := os.Open(servicesYamlFile)

	if err != nil {
		log.Printf("Can not open servces config file %s: %s\n", servicesYamlFile, err)
		return nil, err
	}

	d := yaml.NewDecoder(f)

	cfg := &ServicesConfig{}

	if err := d.Decode(cfg); err != nil {
		log.Printf("Can not parse services config %s: %s\n", servicesYamlFile, err)
		return nil, err
	}

	return cfg, nil
}
