package app

import "context"

type App struct {
	servers *Servers

	errorChan chan error
}

func (app *App) Run(ctx context.Context, cancel context.CancelFunc) error {
	go app.errorHandler(ctx, cancel)

	go app.servers.HTTPServer.Start(app.errorChan)

	return nil
}

func (app *App) Stop() {
	app.servers.Stop()

	close(app.errorChan)
}
