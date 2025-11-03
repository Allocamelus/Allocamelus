qtpl:
	qtc -dir=web/template

run:
	go run ./cmd/allocamelus

generate:
	sqlc generate
	go generate ./...
build-bun:
	cd ./web/app; bun run build;

build-go:
	go build -ldflags="-s -w" -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus

build-go-alpine:
	go build -ldflags="-s -w" -tags=alpine -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus
build-go-alpine-debug:
	go build -gcflags="all=-N -l" -tags=alpine -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus
build-setup:	
	go build -ldflags="-s -w" -o ./cmd/allocamelus-setup/allocamelus-setup ./cmd/allocamelus-setup

build: build-go build-yarn