package app

import "github.com/sirupsen/logrus"

func (app *App) errorHandler() {
	for {
		select {
		case err := <-app.errorChan:
			logrus.Fatal(err)
		case <-app.ctx.Done():
			return
		}
	}
}
