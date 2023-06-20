package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"exchange/cmd/app"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	defer cancel()

	logrus.Info("starting application...")

	app, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	if err = app.Run(ctx, cancel); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("application started =)")

	<-ctx.Done()

	logrus.Info("application stopped.")
}
