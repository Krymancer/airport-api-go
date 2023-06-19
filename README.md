# Airport Go API

This implements in Go the [Airport API Reference](https://github.com/krymancer/airport-api-reference) specification.
The goal of this project is to provide a reference implementation of the Airport API specification in Go that can be used as a starting point for other projects, provide an working example of the API, and to provide a test suite for the specification being able to compare the results of other implementations in differente languages and frameworks.

## Running the API

To run the API you need to have Go installed in your system. You can download it from [here](https://golang.org/dl/).

Once you have Go installed, you can run the API with the following command:

```bash
go run cmd/main.go
```

or

```bash
make server
```

## Technology stack

The API is implemented in Go using the [Gin](github.com/gin-gonic/gin) framework, uses [PostgreSQL](https://www.postgresql.org/) and [Gorm](https://gorm.io/) for database and ORM and uses [Viper](https://github.com/spf13/viper) to manage environment variables.
