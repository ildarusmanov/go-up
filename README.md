# Go-Up

This is a simple package to simplify Golang applications development.

Usage example:

```
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ildarusmanov/go-up/app"
)

func main() {
	log.Println("[+] Starting")

	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

  createApp(ctx)

	log.Println("[*] Started")

	<-sigchan

	cancel()

	time.Sleep(time.Second * 5)

	log.Println("[x] Finished")

	os.Exit(0)
}

type Printer struct{}

func (p *Printer) SayHello(name string) {
	log.Printf("Hello %s!\n", name)
}

func createApp(ctx context.Context) *app.Application {
	a := app.NewApplication(ctx, nil, nil)

	a.SetConfig("userName", "User")

	a.AddServiceFactory("printer", func(ctx context.Context) (app.Service, error) {
		return &Printer{}, nil
	})

	a.AddServiceFactory("hello", func(ctx context.Context) (app.Service, error) {
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Printf("[*] Hello exited with: %s", ctx.Err())
					return
				default:
					p, err := ctx.Value("Application").(*app.Application).GetService("printer")
					if err != nil {
						log.Println(err)
					} else {
						name, _ := ctx.Value("Application").(*app.Application).GetConfig("userName")

						p.(*Printer).SayHello(name)
					}
					time.Sleep(time.Second * 1)
				}
			}
		}()

		return nil, nil
	})

  return a
}

```

See [examples](https://github.com/ildarusmanov/go-up-examples) to learn more.
