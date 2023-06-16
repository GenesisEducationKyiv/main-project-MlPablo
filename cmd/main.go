package main

import (
	"context"
	"os"
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

	ctx, cancel := context.WithCancel(context.Background())

	logrus.Info("starting application...")

	app, err := app.New(ctx, cancel)
	if err != nil {
		logrus.Fatal(err)
	}

	if err = app.Run(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("application started =)")

	go syscallWait(cancel)
	<-ctx.Done()

	logrus.Info("application stopped.")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}
