package main

import (
	"log"
	"os"
	"text/template"
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

	for _, srv := range cfg.Factories {
		createAppFactory(wdir, srv)
	}
}

func createAppFactory(wdir string, srv *Factory) {
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
