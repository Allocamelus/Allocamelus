qtpl:
	qtc -dir=web/template

run:
	go run ./cmd/allocamelus

build-yarn:
	cd ./web/app; yarn run build;

build-go:
	go build -ldflags="-s -w" -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus

build-go-alpine:
	go build -ldflags="-s -w" -tags=alpine -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus

build: build-go build-yarn