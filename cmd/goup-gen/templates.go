package main

const MainTemplate = `
package main

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

	a.Stop()
	cancel()

	time.Sleep(time.Second * 5)

	log.Println("[x] Finished")

	os.Exit(0)
}
`

const AppTemplate = `
package app

import (
	"context"
	"log"

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

  {{range .Factories}}
	if err := a.AddServiceFactory("{{.FactoryServiceName}}", {{.FactoryTypeName}}Factory); err != nil {
		log.Fatal(err)
	}
  {{end}}

	return a
}

func (a *App) Stop() {
	log.Println("[*] Server stopped")
}

`

const FactoryTemplate = `
package app

import (
	"context"
	"os"

	"github.com/ildarusmanov/go-up/goup"
)

func {{.FactoryTypeName}}Factory(ctx context.Context) (interface{}, error) {
	return nil, nil
}

`
