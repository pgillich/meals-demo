# Meals demo

Demo for a meals service. It's written in Go.

## Development

### Prerequisites

* Docker
* `make`
* Golang compiler, linters are not needed (they are run in container)

### Building

Builds exacutable binary to `build/bin`.

```sh
make build
```

### Checks

Includes belows:

* `make test`: `go test`
* `make lint`: <https://github.com/golangci/golangci-lint>
* `make shellcheck`: <https://github.com/koalaman/shellcheck>
* `make mdlint`: <https://github.com/DavidAnson/markdownlint-cli2>

```sh
make check
```
