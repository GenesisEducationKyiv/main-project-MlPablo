package app

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (app *App) errorHandler(ctx context.Context, cancel context.CancelFunc) {
	for {
		select {
		case err := <-app.errorChan:
			cancel()
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
