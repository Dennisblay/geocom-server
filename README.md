# Project: Ordering App server

This is the server side of the ordering app. It is a RESTful API that allows users to create, read, update, and delete
orders. The server is written in Go and uses a Postgres database to store the orders.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes. See deployment for notes on how to deploy the project on a live system.

## Documentation
Visit this link to view the Database Documentation : [DB Documentation]( https://dbdocs.io/dennisboachie9/ordering
)

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```