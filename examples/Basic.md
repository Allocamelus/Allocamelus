# Basic Setup

[Docker-compose](./DockerCompose.md) setup (recommended)

## Requirements

[Golang](https://golang.org/dl/) (Tested with `1.17+`)

Sql database (Recommended: [MariaDB](https://mariadb.com/downloads/))

[Redis](https://redis.io/download) server

## Building

Get the latest version and build:

```sh
git clone https://github.com/Allocamelus/Allocamelus.git
cd Allocamelus
make build
```

The file(s) can be found at:

- Executable: `cmd/allocamelus/allocamelus`
- Web app: `web/app/dist/`

## Use

Setup sql database using [allocamelus.sql](./allocamelus.sql) for tables

Setup Redis see [redis.io/topics/quickstart](https://redis.io/topics/quickstart)

Copy and edit [config.json](./config.json)

### Running

```sh
./allocamelus -c ./path/to/config.json
```
