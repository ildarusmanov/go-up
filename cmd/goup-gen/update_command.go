package main

import (
	"log"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v3"
)

func UpdateCommand(wdir string) {
	createAppServices(wdir)
}

func createAppServices(wdir string) {
	servicesYamlFile := wdir + "/" + GoupConfigFile
	appTmpl := template.Must(template.New("app").Parse(AppTemplate))
	appFile, err := os.Create(wdir + "/app/app.go")

	if err != nil {
		log.Fatalf("Can not create app/app.go file: %s", err)
	}

	cfg, err := parseConfigYaml(servicesYamlFile)
	if err != nil {
		log.Fatalf("Can not parse yaml: %s", err)
	}

	if err := appTmpl.Execute(appFile, cfg); err != nil {
		log.Printf("Can not create app/app.go: %s", err)
	}

	for _, srv := range cfg.Services {
		createAppServiceFactory(wdir, srv)
	}
}

func createAppServiceFactory(wdir string, srv *ServiceFactory) {
	factoryTmpl := template.Must(template.New("factory").Parse(FactoryTemplate))
	factoryName := wdir + "/app/" + srv.FactoryFilename()
	factoryFile, err := os.Create(factoryName)

	if err != nil {
		log.Printf("Can not create factory file %s: %s\n", factoryName, err)
		return
	}

	if err := factoryTmpl.Execute(factoryFile, srv); err != nil {
		log.Printf("Can not create factory %s: %s", factoryName, err)
	}

}

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
