# goservice

Base web service project in Go, currently aiming for basic http api implementation.

End goal is to have a single project that could be run using either gRPC or REST api.

## Requirements

- Go v1.16
- make

## Environment Variables

Service:

- `API_PORT`: 80
- `LOG_LEVEL`: info

Docker:

- `DOCKER_REPO`: user/repo

## Usage

Only need to clone the repo and run

```bash
make
```

Or run with Docker, you have to set the Docker env vars.

```bash
make run-docker
```

## License

[MIT](https://github.com/sergionunezgo/gorest/LICENSE)
