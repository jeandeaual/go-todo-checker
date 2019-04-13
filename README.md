# TODO Checker

Command-line tool to create a list of TODOs written within Go source code.

## Prerequisites

Go 1.6 or later.

## Build Instructions

Make sure the repository is located in `GOPATH`, and run `make` or:

```
go build ./cmd/todo
```

## Usage

```
Usage: todo [PACKAGE]

Flags:
  -pattern string
        Pattern to search for in the package comments (only used without -server) (default "TODO")
  -port int
        Server port number (only used with -server) (default 80)
  -server
        Run the program in server mode
```

### Examples

Command-line:

```
$ ./todo fmt
/usr/lib/go/src/fmt/scan.go:732:
TODO: accept N and Ni independently?
```

Server mode:

* Run the server

  ```
  $ ./todo -server -port 8080
  2006/01/02 15:04:05 Listening on port 8080
  ```

* Query the API:

  ```
  $ curl -w "\n" http://localhost:8080/fmt
  [{"filename":"/usr/lib/go/src/fmt/scan.go","line":732,"text":"TODO: accept N and Ni independently?\n"}]

  $ curl -w "\n" http://localhost:8080/net/http/httptest
  []

  $ curl -w "\n" http://localhost:8080/bytes?pattern=FIXME
  [{"filename":"/usr/lib/go/src/bytes/buffer.go","line":25,"text":"FIXME: it would be advisable to align Buffer to cachelines to avoid false\nsharing.\n"}]
  ```

## Testing

Make sure testify is installed:

```
go get github.com/stretchr/testify
```

Run the tests using `make test` or:

```
go test -v ./...
```

To generate a code coverage report using `make coverage` or:

```
go test -coverpkg=./... -coverprofile="coverage.out" ./...
go tool cover -html="coverage.out"
```

## API Documentation

The documentation of the HTTP API (when using the `-server` flag) can be
generated using `make doc`, or by following these instructions:

* Install [swag](https://github.com/swaggo/swag)
  * `go get -u github.com/swaggo/swag/cmd/swag`
  * Make sure `$GOPATH/bin` is in your `$PATH`
* Run the tool

  ```
  swag init -g ./cmd/todo/api.go
  ```

This will generate `docs/swagger.yaml` and `docs/swagger.json`, which can be
viewed using the [Swagger editor](https://editor.swagger.io/) for example.

The documentation is also available as an [OpenAPI specification file](docs/openapi.yaml),
which can be viewed using the same tool as the auto-generated file.
