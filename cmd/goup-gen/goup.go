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

	wdir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Can not detect current directory: %s", err)
	}

	CreateAppDir(wdir)
	CreateAppServices(wdir)
	CreateMainFile(wdir)

	log.Println("Done!")
}

func CreateAppDir(wdir string) {
	if err := os.Mkdir(wdir+"/app", os.FileMode(0777)); err != nil {
		log.Fatalf("Can not create app/ directory: %s", err)
	}
}

func CreateAppServices(wdir string) {
	appTmpl := template.Must(template.New("app").Parse(AppTemplate))
	appFile, err := os.Create(wdir + "/app/app.go")

	if err != nil {
		log.Fatalf("Can not create app/app.go file: %s", err)
	}

	if *servicesYamlFile != "" {
		cfg, err := ParseConfigYaml(*servicesYamlFile)
		if err != nil {
			log.Fatalf("Can not parse yaml: %s", err)
		}

		if err := appTmpl.Execute(appFile, cfg); err != nil {
			log.Printf("Can not create app/app.go: %s", err)
		}

		for _, srv := range cfg.Services {
			CreateAppService(wdir, srv)
		}
	}
}

func CreateAppService(wdir string, srv *ServiceFactory) {
	factoryTmpl := template.Must(template.New("factory").Parse(FactoryTemplate))
	srvFName := wdir + "/app/" + srv.Filename
	srvFile, err := os.Create(srvFName)

	if err != nil {
		log.Printf("Can not create service file %s: %s\n", srvFName, err)
		return
	}

	if err := factoryTmpl.Execute(srvFile, srv); err != nil {
		log.Printf("Can not create factory %s: %s", srvFName, err)
	}
}

func CreateMainFile(wdir string) {
	mainTmpl := template.Must(template.New("main").Parse(MainTemplate))
	mainFile, err := os.Create(wdir + "/main.go")

	if err != nil {
		log.Fatalf("Can not create main.go file: %s", err)
	}

	if err := mainTmpl.Execute(mainFile, map[string]string{"Pkgname": *pkgname}); err != nil {
		log.Fatalf("Can not create main.go: %s", err)
	}
}

func ParseConfigYaml(servicesYamlFile string) (*ServicesConfig, error) {
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
