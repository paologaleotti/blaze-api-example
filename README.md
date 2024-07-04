# blaze example app

This repo contains an example of a HTTP API built using the [blaze](https://github.com/paologaleotti/blaze) template.

It features a simple CRUD API for managing todos, including:

- A basic blaze scaffolding (including logging, configuration, structure...)
- Clean structure and dependency inversion
- SQL data storage using SQLite and [sqlx](https://github.com/jmoiron/sqlx)
- Database migrations using [goose](https://github.com/pressly/goose)

## Running the application

Requirements:
- Go >= 1.22
- sqlite3
- goose

Before running, export the required environment variables for the service (see `internal/api/env.go`) and for Goose (see `.env.template` for all needed env).

To run the app, first create the database and apply the migrations:

```sh
goose up
```

Then build and run the API:

```sh
make
```

```sh
./bin/api/main
```

Done! The API will be available at `http://localhost:3000`.
