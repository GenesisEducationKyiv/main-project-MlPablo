package app

import "context"

type App struct {
	servers *Servers

	errorChan  chan error
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (app *App) Run() error {
	go app.errorHandler()

	go app.servers.HTTPServer.Start(app.errorChan)

	return nil
}

func (app *App) Stop() {
	app.servers.Stop()

	close(app.errorChan)
}
