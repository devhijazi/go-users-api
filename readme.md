# Go API with Kafka

Complete user API with PostgreSQL database and Kafka for event streaming

## TODO

- [x] Implement Refresh Session and Token
- [ ] Fix Delete method
- [ ] Implement Kafka for mail sender
- [x] Docs
- [x] Improve Docker configuration
- [x] Fix token session
- [x] Implement Pagination

## Where can this api scale to ?

- Create one more entity and role for multiple sessions.

## Tecnologies

- [Go](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [Swagger](https://swagger.io/)

### Gin

Gin is a web framework written in Golang.

### Swagger

Swagger allows you to describe the route structure of your API.

### Kafka

Kafka enables orchestration of event streaming, allowing for automatic and high-performance management of message sending queues.

## Getting Started

### Creating the Postgres Container

First, you need to create the Postgres container. To do this, run the following command:

```sh
$ docker-compose up postgres -d
```

**Error Tratatives Map**

```diff

# HTTP Errors
+ 0 - 999

# Functional Errors
+ 1000 - 1999

# Entities Error
+ 2000 - 3999
```

### Email Service

Will be implemented...

### Configuring Environment Variables

Copy the example variables from the `.env.example` file to `.env`.

**Environment Variables**

```diff
# API
+ PORT=
+ URL=

# JWT
+ JWT_TOKEN=

# KAFKA (Not Necessary Yet)
- KAFKA_BROKERS=
- KAFKA_CLIENT_ID=
- KAFKA_TOPIC_ISSUE_EMAIL=
- KAFKA_TOPIC_EMAIL_RESPONSE=

# DATABASE
+ PG_USER=
+ PG_PASSWORD=
+ PG_HOST=
+ PG_PORT=
+ PG_DATABASE=
```

### Installing Dependencies

To do this, run the following command:

```sh
$ go mod download
```

### Starting the API

After completing the above steps, start the API by running the following command:

```sh
$ go run cmd/server/main.go
```

Or

> Lembre-se de ter instalado a CLI do AIR. <a href="https://github.com/cosmtrek/air" target="_blank">Link do repositório</a>.

```sh
$ air
```

### Generating Documentation

If you make any changes to the Swagger comments, you need to run the `go-swagger` command to generate the documentation files. To do this, run the following command:

_Linux_

```sh
$ scripts/build_docs.sh
```

_Windows_

```sh
$ bash .\scripts\build_docs.sh
```
