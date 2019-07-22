package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	yaml "gopkg.in/yaml.v3"
)

func InitCommand(wdir string) {
	pname := flag.Arg(1)

	if pname == "" {
		log.Printf("Invalid project name")
		return
	}

	pkg := flag.Arg(2)

	if pkg == "" {
		log.Printf("Invalid project name")
		return
	}

	pdir := wdir + "/" + pname
	dirs := []string{pdir, pdir + "/app", pdir + "/services"}

	for _, dir := range dirs {
		if err := os.Mkdir(dir, os.FileMode(0777)); err != nil {
			log.Fatalf("Can not create %s: %s", dir, err)
		}

		log.Printf("New directory %s successfully created", dir)
	}

	fpath := pdir + "/" + GoupConfigFile

	cfgFile, err := os.Create(fpath)

	if err != nil {
		log.Fatalf("Can not create %s: %s", fpath, err)
	}

	defer cfgFile.Close()

	log.Printf("New file %s successfully created", fpath)

	cfg := NewServicesConfig(pkg)
	enc := yaml.NewEncoder(cfgFile)

	defer enc.Close()

	if err := enc.Encode(cfg); err != nil {
		log.Fatalf("Can not write config file: %s", err)
	}

	mainTmpl := template.Must(template.New("main").Parse(MainTemplate))
	mainFile, err := os.Create(pdir + "/main.go")

	if err != nil {
		log.Fatalf("Can not create main.go file: %s", err)
	}

	if err := mainTmpl.Execute(mainFile, map[string]string{"Pkgname": pkg}); err != nil {
		log.Fatalf("Can not create main.go: %s", err)
	}

	log.Println("Application initialized")
}
