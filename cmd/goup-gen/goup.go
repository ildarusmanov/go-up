package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var (
	pkgname = flag.String("pkgname", "", "Your application package name.")
)

func main() {
	flag.Parse()

	mainTmpl := template.Must(template.New("main").Parse(MainTemplate))
	appTmpl := template.Must(template.New("app").Parse(AppTemplate))
	// factoryTmpl := template.Must(template.New("factory").Parse(FactoryTemplate))

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
	appTmpl.Execute(appFile, map[string]interface{}{})

	log.Println("Done!")
}
