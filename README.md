# TODO Checker

Command-line tool to create a list of TODOs written within Go source code.

## Prerequisites

Go 1.7 or later.

## Build Instructions

```
go build -o todo
```

## Usage

```
$ ./todo fmt
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
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
