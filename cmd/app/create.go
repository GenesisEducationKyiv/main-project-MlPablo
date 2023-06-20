package app

import (
	"context"
)

func New(ctx context.Context, cancelFunc func()) (*App, error) {
	errChan := make(chan error)

	servers, err := createServers()
	if err != nil {
		return nil, err
	}

	if err = createServicesAndHandlers(servers); err != nil {
		return nil, err
	}

	return &App{
		servers:    servers,
		errorChan:  errChan,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}, nil
}
