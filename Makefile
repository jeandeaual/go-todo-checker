.PHONY: build
build:
	@go build ./cmd/todo

.PHONY: test
test:
	@go test -v ./...

.PHONY: coverage
coverage:
	@go test -coverpkg=./... -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

.PHONY: doc
doc:
	@swag -v || go get -u github.com/swaggo/swag/cmd/swag
	@swag init -g ./cmd/todo/api.go
	@rm -f ./docs/docs.go
