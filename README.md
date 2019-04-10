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
Usage: test.exe [PACKAGE]

Flags:
  -pattern string
        Pattern to search for in the package comments (default "TODO")
  -server
        Server mode
```

Example:

```
$ ./todo fmt
/usr/lib/go/src/fmt/scan.go:732:
TODO: accept N and Ni independently?
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
