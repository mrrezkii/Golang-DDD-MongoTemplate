package main

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/di"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/infrastructure"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	container := di.Container

	err := container.Invoke(func(sh shared.Holder, inh infrastructure.Holder) error {
		app := cli.NewApp()

		app.Commands = append(app.Commands, NewRunCommand(sh, inh)...)

		// - main-command-end

		if err := app.Run(os.Args); err != nil {
			log.Printf("%s", err)
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func NewRunCommand(sh shared.Holder, inh infrastructure.Holder) []cli.Command {
	return []cli.Command{
		{
			Name:  "run",
			Usage: "run as service",
			Action: func(_ *cli.Context) error {
				// - graceful shutdown
				sig := make(chan os.Signal)
				signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

				go inh.ListenHTTP()

				log.Println("application started!")

				<-sig

				// - graceful shutdown here
				sh.Close()

				return nil
			},
		},
	}
}
