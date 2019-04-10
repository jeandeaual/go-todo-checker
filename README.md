# TODO Checker

Command-line tool to create a list of TODOs written within Go source code.

## Prerequisites

Go 1.6 or later.

## Build Instructions

```
go build -o todo
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

Run the tests:

```
go test
```

## API Documentation

The documentation of the HTTP API (when using the `-server` flag) is written
in the [OpenAPI specification file](openapi.yaml).\
It can be viewed using the [Swagger editor](https://editor.swagger.io/) for
example.
