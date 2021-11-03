qtpl:
	qtc -dir=web/template

tscriptify:
	go run ./scripts

run:
	go run ./cmd/allocamelus

build-npm:
	cd ./web/app; npm run build;

build-go:
	go build -ldflags="-s -w" -o ./cmd/allocamelus/allocamelus ./cmd/allocamelus

build: build-go build-npm