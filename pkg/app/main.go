package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sergionunezgo/go-reuse/pkg/logger"
	"github.com/sergionunezgo/go-reuse/pkg/service"
	"github.com/urfave/cli"
)

var (
	// Reference to the api service, it has to implement io.Closer interface for clean-up.
	serviceRef service.Service
	interrupt  chan os.Signal
)

// Create loads env variables, calls the initService method for setup and starts the service.
func Create(flags []cli.Flag, initService func(cfg *service.Config) (service.Service, error)) *cli.App {
	setupInterruptCloseHandler()

	config := &service.Config{}
	baseFlags := []cli.Flag{
		cli.IntFlag{
			Name:        "api_port",
			EnvVar:      "API_PORT",
			Value:       80,
			Usage:       "port for the web service",
			Destination: &config.Port,
		},
		cli.StringFlag{
			Name:        "log_level",
			EnvVar:      "LOG_LEVEL",
			Value:       "info",
			Usage:       "log level for the logger",
			Destination: &config.LogLevel,
		},
	}

	app := cli.NewApp()
	app.Version = "0.0.0"
	app.Flags = append(baseFlags, flags...)

	app.Action = func(ctx *cli.Context) error {
		err := logger.UseZapLogger(config.LogLevel)
		if err != nil {
			return errors.Wrap(err, "logger UseZapLogger")
		}

		logger.Log.Info("calling initService function")
		service, err := initService(config)
		if err != nil {
			return errors.Wrap(err, "initService")
		}

		serviceRef = service
		return service.Start()
	}

	return app
}

// Interrupt can be used to send an interrupt signal to the running service like syscall.SIGTERM
func Interrupt(sig os.Signal) {
	interrupt <- sig
}

// setupInterruptCloseHandler run a goroutine to listen for interruption signals to perform clean-up.
func setupInterruptCloseHandler() {
	interruptions := make(chan os.Signal, 2)
	interrupt = interruptions
	signal.Notify(interruptions, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptions
		logger.Log.Warn("interruption signal received, starting clean-up")
		if serviceRef != nil {
			serviceRef.Close()
		}
		logger.CloseLogger()
		os.Exit(0)
	}()
}
