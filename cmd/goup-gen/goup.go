package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	pkgname          = flag.String("pkgname", "", "Your application package name.")
	servicesYamlFile = flag.String("cfg", "", "Services config file in YAML format")
)

func main() {
	flag.Parse()

	mainTmpl := template.Must(template.New("main").Parse(MainTemplate))
	appTmpl := template.Must(template.New("app").Parse(AppTemplate))
	factoryTmpl := template.Must(template.New("factory").Parse(FactoryTemplate))

	wdir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Can not detect current directory: %s", err)
	}

	if err := os.Mkdir("app", os.FileMode(0777)); err != nil {
		log.Fatalf("Can not create app/ directory: %s", err)
	}

	mainFile, err := os.Create(wdir + "/main.go")

	if err != nil {
		log.Fatalf("Can not create main.go file: %s", err)
	}

	appFile, err := os.Create(wdir + "/app/app.go")

	if err != nil {
		log.Fatalf("Can not create app/app.go file: %s", err)
	}

	mainTmpl.Execute(mainFile, map[string]string{
		"Pkgname": *pkgname,
	})

	if *servicesYamlFile != "" {
		f, err := os.Open(*servicesYamlFile)

		if err != nil {
			log.Fatalf("Can not open servces config file %s: %s", *servicesYamlFile, err)
		}

		d := yaml.NewDecoder(f)

		cfg := &ServicesConfig{}

		if err := d.Decode(cfg); err != nil {
			log.Fatalf("Can not parse services config %s: %s", *servicesYamlFile, err)
		}

		appTmpl.Execute(appFile, cfg)

		for _, srv := range cfg.Services {
			srvFName := wdir + "/app/" + srv.Filename
			srvFile, err := os.Create(srvFName)

			if err != nil {
				log.Printf("Can not create service file %s: %s\n", srvFName, err)
				continue
			}

			factoryTmpl.Execute(srvFile, srv)
		}
	}

	log.Println("Done!")
}
