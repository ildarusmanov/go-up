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

	a := app.NewApp(ctx)
	log.Println("[*] Started")

	<-sigchan
	cancel()

	time.Sleep(time.Second * 5)
	log.Println("[x] Finished")
}
`

const AppTemplate = `package app

import (
	"context"
	"log"

{{range .Services}}
  "{{.ServicePackage.Import}}"
{{end}}
	"github.com/ildarusmanov/go-up/config"
	"github.com/ildarusmanov/go-up/goup"
)

type App struct {
	*goup.Application
}

func NewApp(ctx context.Context) *App {
	a := &App{Application: goup.NewApplication()}

	a.WithContext(ctx).WithConfig(config.NewEnvConfig())

	requiredConfigKeys := []string{}

	if err := a.Config().RequireKeys(requiredConfigKeys); err != nil {
		log.Fatal(err)
	}

	// create services
  {
{{range .Services}}
		c{{.FactoryName}}, err := {{.FactoryName}}FactoryConfigGetter(a.Config())

		if err != nil {
			log.Fatal(err)
		}

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
	go func(){
		for {
			select {
			case <-ctx.Done():
				log.Print("[x] Stop application")
				// add services handlers
				// e.g.: someService.Close()
				// ...
				return
			default:
				continue;
			}
		}
	}()

	return a
}
`

const FactoryTemplate = `package app

import (
	"context"

	"{{.ServicePackage.Import}}"
{{range .Dependencies}}
  "{{.DependencyPackage.Import}}"
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
	if fcfg, err := {{.FactoryName}}FactoryConfigGetter(cfg goup.ConfigManager); err != nil {
		return nil, err
	}

	// create service with type {{.ServiceType}}

	return nil, nil
}

`
