# API

[![Go Report Card](https://goreportcard.com/badge/github.com/syntaqx/api)](https://goreportcard.com/report/github.com/syntaqx/api)
[![Docker Image Size](https://img.shields.io/docker/image-size/syntaqx/api)](https://hub.docker.com/r/syntaqx/api)

My Personal API for things I want to API.

## Getting started

```sh
git clone git@github.com:syntaqx/api.git && cd "$_"
```

## Development

```sh
docker compose up -d --build
```

> [!NOTE]
> To bind to host ports, use the Docker Compose Overrides and adjust them to
> your liking.
> `cp compose.override.example.yml compose.override.yml`
