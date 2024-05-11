SHELL := /bin/bash

init:
	go mod init

tidy:
	go mod tidy

run-migration:
	go run migration/migration.go

run:
	go run main.go

build:
	go build

mockgen:
	rm -r repo/mock
	mkdir repo/mock
	mockgen -source=repo/repo.go -destination=repo/mock/mock_repo.go -package=repo_mock

test:
	go test ./handler/...

coverage-report:
	go test -coverprofile cover.out ./handler/...
	go tool cover -html cover.out

mock-gen:
