# BTC Application

This repository contains a Golang server that implements an API for retrieving the current BTC to UAH exchange rate, managing email subscriptions for rate change notifications, and sending rate to subscribed users. App is using <https://currencyapi.com/> for getting BTC/UAH rate.

## API Specification

The API follows the Swagger 2.0 specification. You can find the detailed API documentation in the api/openapi/swagger.yaml file.

### Endpoints

- `/rate`: Retrieves the current BTC to UAH exchange rate.
- `/subscribe`: Subscribes an email address to receive rate notifications.
- `/sendEmails`: Sends rate BTC to UAH to all subscribed email addresses.

Please note that all API endpoints are prefixed with `/api` to ensure a consistent path structure.

## Requirements

- Golang
- Docker

## Getting Started

### Manually

1. Clone the repository:

```shell
git clone https://github.com/MlPablo/exchange.git
cd exchange
```

2. Build the Docker image:
```shell
docker build -t exchange .
```

3. Run the Docker container:
```shell
docker run -p 8080:8080 exchange
```

4. Or run locally:
```shell
go run ./...
```

### Using make

- To run the program with the binary, use the following command:
  ```
  make run
  ```
  This will compile the project and execute the resulting binary.
  
- To compile and run the Docker container, use the following command:
  ```
  make docker
  ```
  This will build a Docker image for the application and run it in a container.

> **_NOTE:_**: Make sure you have Docker installed and running on your machine.

Your server will start on port 8080

### Project structure
The project follows a modular structure with different directories serving specific purposes. Here's a breakdown of the directory structure:
```shell
/cmd/web/main.go         - Entry point of the application where all modules are set up.

/pkg
    /domain              - Contains data models and their communication API.
    /http                - Handles routes and has access to services.
    /infrastructure      - Handles third-party API logic such as sending emails and retrieving currency information. These modules must implement the interfaces defined by the services that use them.
    /repository          - Contains logic for storing data and interacts only with the domain.
    /services            - Implements the business logic of the application. It depends on the domain and repository modules (and sometimes other services).

.env                     - Contains environment variables where you can define server port, email user, and other configuration options.
```

The project structure is organized in a way that promotes modularity, separation of concerns, and easy navigation. Each directory has a specific purpose and encapsulates related functionality, allowing for better maintainability and scalability of the application.
