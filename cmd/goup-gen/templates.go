package main

const MainTemplate = `package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.Pkgname}}/app"
)

func main() {
	log.Println("[+] Starting")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	stopApplication := app.StartApplication(ctx)
	log.Println("[*] Started")

	<-sigchan
	cancel()
	stopApplication()

	time.Sleep(time.Second * 5)
	log.Println("[x] Finished")
}
`

const AppTemplate = `package app

import (
	"context"
	"log"

{{range .Factories}}
  {{.ServicePackage.GetDefinition}}
{{end}}
	"github.com/ildarusmanov/go-up/config"
	"github.com/ildarusmanov/go-up/goup"
)

func StartApplication(ctx context.Context) goup.StopApplicationHandler {
	cfg := config.NewEnvConfig()

	requiredConfigKeys := []string{}

	if err := cfg.RequireKeys(requiredConfigKeys); err != nil {
		log.Fatal(err)
	}

	// create services
  {
{{range .Factories}}
		s{{.FactoryName}}, err := {{.FactoryName}}Factory(
			ctx,
			c,{{range .Dependencies}}
			s{{.FactoryName}},{{end}}
		)

		if err != nil {
			log.Fatal(err)
		}
{{end}}
  }

	// stop the application
	return func(){
		log.Println("[x] Terminating...")
		// add code here
		log.Println("[x] Done")
	}
}
`

const FactoryTemplate = `package app

import (
	"context"

	{{.ServicePackage.GetDefinition}}
{{range .Dependencies}}
  {{.DependencyPackage.GetDefinition}}
{{end}}
	"github.com/ildarusmanov/go-up/goup"
)

type {{.FactoryName}}FactoryConfig struct{}

func {{.FactoryName}}FactoryConfigGetter(cfg goup.ConfigManager) ({{.FactoryName}}FactoryConfig, error) {
	return &{{.FactoryName}}FactoryConfig{}, nil
}

func {{.FactoryName}}Factory(
	ctx context.Context,
	cfg goup.ConfigManager,{{range .Dependencies}}
	s{{.FactoryName}} {{.Type}},{{end}}
) ({{.ServiceType}}, error) {
	if fcfg, err := {{.FactoryName}}FactoryConfigGetter(cfg); err != nil {
		return nil, err
	}

	// create service with type {{.ServiceType}}

	return {{.ServiceConstructor.Signature}}
}

`
