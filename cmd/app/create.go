package app

func New() (*App, error) {
	errChan := make(chan error)

	servers, err := createServers()
	if err != nil {
		return nil, err
	}

	services, err := createServices()
	if err != nil {
		return nil, err
	}

	registerHandlers(servers, services)

	return &App{
		servers:   servers,
		errorChan: errChan,
	}, nil
}
