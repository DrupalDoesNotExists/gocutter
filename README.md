# Go CookieCutter Template

This is a cookiecutter template for the Golang microservices. It solves a lot of boilerplate related issues and makes a great experience of productive work.

## Features

* Dockerfile with multistage build
* Docker Compose for convenient development
* Well-defined file and folder structure
* Documented Makefile with lots of useful targets
* Ready-to-go Postgres support
* Built on top of Gin, SqlX, PGX, Zap, GoJWT and others
* Out of the box JWT support with mapped marshaller
* Golang Migrate tool support (even in Compose)
* Protobuf (with code-gen) support
* Easy to use wizard (even license choose included)

## Directory structure

```
├── deployments
│   ├── docker-compose.yaml
│   └── local.env
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   └── controllers
├── LICENSE
├── main.go
├── Makefile
├── migrations
├── proto
├── resources
│   ├── jwtkey
│   └── jwtkey.pub
└── tests
```

Every directory contains some useful README that describes all the needed stuff.
