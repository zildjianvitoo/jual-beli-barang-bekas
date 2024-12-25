# Requirements

- PostgreSQL
- Go >= ^1.23.4
- use air library to run this project [air github & docs](https://github.com/air-verse/air)

# How to run

1. Copy the `.env.example` to `.env` and adjust the value
2. Open your docker desktop app
3. Open your terminal, change directory to this project, and run "docker compose up -d"
4. Run the http server with `air`

# The interesting parts in this app

To make this app scalable, I built the backend system for a second-hand goods marketplace with a focus on achieving low latencies. To achieve this, I optimized the application to ensure it handles high traffic efficiently. However, due to time constraints, there are still some parts I couldnt complete within two weeks.

<!-- # API Documentation

Please find the ``, and import it to your postman
to test the API. -->

# Directory Structure

This app follows the Uncle Bob Clean Architecture in terms of the dependencies rules.
This spec allows the app to be more easy to maintain because of the dependencies rules.
This app also following the port-adapter pattern as you can see we can implement
the entire application by just implementing the transport layer whether it's for
the incoming or outbound connection such as switching the database or switching the
tranport for receiving requests.

```
jual-beli-barang-bekas
├── config                  # Configuration and setup env variables
├── internal
│   ├── api
│   │   ├── graphql         # GraphQL handlers (not used)
│   │   ├── grpc            # gRPC handlers (not used)
│   │   └── rest            # REST API handlers
│   │       ├── handlers    # HTTP handlers/controllers
│   │       ├── response    # Response formatting
│   │       └── server      # REST API server
│   ├── domain              # Domain layer
│   ├── dto                 # Data Transfer Objects (Entities and domain models)
│   ├── helper              # Utility functions and helpers
│   ├── service             # Domain-specific services
│   ├── repository          # Persistence layer
│   └── service             # Application services
├── pkg                     # Shared utilities and packages
```

## server.go file

Infrastructure layer is a gateway to external world of the application, such as
connecting to rest api, postgresql, etc. All outbound connection should use this
as a gateway

## Domain directory

Domain is used to put all of enterprise and business rules. This is the most inward layer
in terms of clean architecture by Uncle Bob.
All of the business rules should be put here.
I'm also using domain service terms which is comes from DDD (domain design driven)
which is used to put business logic that not fit into a domain entity

## Repository directory

This is the concrete implementations of domain repository.
This should be specific and can be dependent and coupled to other package
(such as postgres, redis (for future improvement), etc)

## Config directory

This directory contains application config

## Utils directory

Contains all of helpers and utilities used by the application
