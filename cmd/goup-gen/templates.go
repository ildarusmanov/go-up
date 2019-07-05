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

	a.Stop()
	cancel()

	time.Sleep(time.Second * 5)

	log.Println("[x] Finished")

	os.Exit(0)
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
{{range .Services}}
	if err := a.AddServiceFactory("{{.ServiceName}}", {{.FactoryName}}Factory); err != nil {
		log.Fatal(err)
	}
{{end}}
	return a
}
{{range .Services}}
func (a *App) {{.MethodName}}() (*{{.ServicePackage.Name}}.{{.ServiceType}}, error) {
	s, err := a.GetService("{{.ServiceName}}")

	if err != nil {
		return nil, err
	}

	srv, ok := s.(*{{.ServicePackage.Name}}.{{.ServiceType}})

	if !ok {
		return nil, errors.New("Incorrect service type")
	}

	return srv, nil
}
{{end}}
func (a *App) Stop() {
	log.Println("[*] Server stopped")
}

`

const FactoryTemplate = `package app

import (
	"context"

	"{{.ServicePackage.Import}}"
	"github.com/ildarusmanov/go-up/goup"
)

func {{.FactoryName}}Factory(ctx context.Context) (interface{}, error) {
	return {{.ServicePackage.Name}}.New{{.ServiceType}}()
}

`
