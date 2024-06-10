# API

[![Docker Image Size](https://img.shields.io/docker/image-size/syntaqx/api)](https://hub.docker.com/r/syntaqx/api)
[![Go Report Card](https://goreportcard.com/badge/github.com/syntaqx/api)](https://goreportcard.com/report/github.com/syntaqx/api)
[![codecov](https://codecov.io/gh/syntaqx/api/graph/badge.svg?token=M5iaJ6FseZ)](https://codecov.io/gh/syntaqx/api)

My Personal API for things I want to API.

## Getting started

### Clone the repository

```sh
git clone git@github.com:syntaqx/api && cd "$(basename "$_")"
```

### Starting the environment

```sh
docker compose up -d --build
```

> [!NOTE]
> You may prefer to develop using `go` directly. You can reference the `Dockerfile`
> for necessary steps to run the code on your system, but given the steps will be
> different for each host operating system and their versions, this route will
> remain undocumented.

#### Binding to `localhost`

In order to access the container ports over `localhost` you'll need to override the default
`compose.yml` port values and specify your preferred host port. To use the default
values, or quickly populate the file, simply:

```sh
cp compose.override.example.yml compose.override.yml
```

> [!NOTE]
> Any changes made to the `compose.override.yml` will be ignored by git, to feel free to
> use this to make any environment-specific changes or overrides without affecting others
> local settings.

Then (re)start your containers with the new port bindings:

```sh
docker compose up -d
```
