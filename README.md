# Logistics Manager App (Technical Test)

## Getting Started

To get started with this application, you need to have Go, Docker, and Docker Compose installed on your system.

### Prerequisites

- Go: [Installation instructions](https://golang.org/doc/install)
- Docker: [Installation instructions](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation instructions](https://docs.docker.com/compose/install/)

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/hesam-khorshidi/order-management/
    ```

2. Change into the project directory:

    ```bash
    cd order-management
    ```

### Running Locally

To run the application locally, follow these steps:

1. Start the application using `docker-compose`:

    ```bash
    docker-compose up
    ```

2. The application will start, and you can access it at `http://localhost:8080`.

## Project Structure

```
project
└───common
│   └─── utils.go # Contains utility functions
│      
└───controller
│   └─── order_controller.go
|   └─── provider_controller.go
|
└───entity
|   └─── customer.go
|   └─── dto.go
|   └─── order.go
|   └─── provider.go
│      
└───providers
│   └─── database_provider.go
|   
└───repositoty
|   └─── order_repository.go
|   └─── provider_repository.go
│      
└───routers
│   └─── router.gp
|
└───service
|   └─── order_service.go
|   └─── provider_service.go
|   └─── sms_service.go
|
│   README.md
│   Dockerfile # Dockerfile for building the Go application
│   docker-compose.yml # Docker Compose file for running the application and database
│   go.mod # Go module file
│   go.sum # Go module sum file
│   main.go # Main application entry point
│   .gitignore
└   .env # Application configuration
```

## Dependencies

This application uses the following dependencies:

- [Gin](https://github.com/gin-gonic/gin): Web framework for Go.
- [GORM](https://gorm.io/): Object-relational mapping library for Go.

## TODO
Adding swagger to project is a priority to create a better API documentation
