# Setup

Marketplace REST API can be built as a Docker container. To build the server only docker and docker-compose are needed.

## Build and run the server in a Docker container

```sh
git clone https://github.com/pawelktk/marketplace.git
cd marketplace
docker compose up --build
```

## Storage

The database is stored persistently in "pgdata" Docker volume. To remove the volume:

```sh
docker compose down
docker volume rm marketplace_pgdata
```

# API Example Usage

```sh
curl --location 'http://localhost:8080/api/products'
```

# Frontend

Simple frontend server is available at _http://localhost:8080/_
