qtpl:
	qtc -dir=web/template

tscriptify:
	go run ./scripts

run:
	go run ./cmd/allocamelus

build-npm:
	cd ./web/app; npm run build;

build-go:
	cd ./cmd/allocamelus; go build -ldflags="-s -w" -o ./application.so .

build: build-go build-npm