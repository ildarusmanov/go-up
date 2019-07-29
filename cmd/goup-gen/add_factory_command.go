package main

import (
	"flag"
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func AddFactoryCommand(wdir string) {
	servicesYamlFile := wdir + "/" + GoupConfigFile

	cfg, err := parseConfigYaml(servicesYamlFile)
	if err != nil {
		log.Fatalf("Can not parse yaml: %s", err)
	}

	factoryName := flag.Arg(1)
	serviceType := flag.Arg(2)

	cfg.AddFactory(&Factory{
		FactoryName: factoryName,
		ServiceType: serviceType,
	})

	cfgFile, err := os.Create(servicesYamlFile)

	if err != nil {
		log.Fatalf("Can not create %s: %s", servicesYamlFile, err)
	}

	defer cfgFile.Close()

	log.Printf("New file %s successfully created", servicesYamlFile)

	enc := yaml.NewEncoder(cfgFile)

	defer enc.Close()

	if err := enc.Encode(cfg); err != nil {
		log.Fatalf("Can not write config file: %s", err)
	}

	log.Println("Done!")
}
