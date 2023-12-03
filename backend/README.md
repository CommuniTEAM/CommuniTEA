# CommuniTEA Go API Service

## Table of Contents

- [CommuniTEA Go API Service](#communitea-go-api-service)
  - [Table of Contents](#table-of-contents)
- [Development](#development)
  - [Working with the Database](#working-with-the-database)
    - [Database Migrations](#database-migrations)
    - [Go and the Database](#go-and-the-database)

# Development

To set up your local development environment, please have [Docker](https://www.docker.com/products/docker-desktop/) and [golangci-lint](https://golangci-lint.run/) installed on your local machine. Detailed installation instructions for golangci-lint can be found in [the CommuniTEA README](https://github.com/CommuniTEAM/CommuniTEA/tree/main#go-format).

## Working with the Database

Our API service uses a PostgreSQL database housed in a Docker container, which you can stand up by navigating to the project's root directory in your terminal and running:

```
docker volume create communitea-db
docker compose build
docker compose up
```

This will start the entire application, including the API service container and the frontend application container in addition to the database.

### Database Migrations

After creating the docker volume for the project or after pulling changes in the `db/migrations` directory, migrations must be ran against the database in order for the changes to be applied locally. [Dbmate](https://github.com/amacneil/dbmate) is used to manage migrations.

To run existing migrations against the database, navigate to `CommuniTEA/backend` in your terminal and run:

```
docker run --rm -it --network=host -v "$(pwd)/db:/db" ghcr.io/amacneil/dbmate -u "postgres://admin:secret@localhost:15432/communitea-db?sslmode=disable" up
```

To create a new migration file, run:

```
docker run --rm -it --network=host -v "$(pwd)/db:/db" ghcr.io/amacneil/dbmate new <NAME OF YOUR MIGRATION (ex. create_users_tables)>
```

For Dbmate to run migrations successfully, the SQL file must have both `-- migrate:up` (above any creation statements) and `-- migrate:down` (above any deletion statements) present, even if one is left blank.

### Go and the Database

The API service utilizes [sqlc](https://sqlc.dev/) to blend the SQL schema and queries with the Go ecosystem. All queries are written directly in SQL (see `/db/queries`) and sqlc is used to auto-generate the corresponding Go code.

If you make changes to the migrations or the queries, you will need to instruct sqlc to re-generate the corresponding Go code. This is done automatically every time the `goapi` docker container starts, but it can also be done manually by connecting to the container's terminal and running `sqlc generate`.
