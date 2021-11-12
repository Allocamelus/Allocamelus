# Docker-compose Setup

[Basic](./Basic.md) setup

## Requirements

[docker/engine](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/)

## Setup

### Building

Get the latest version and build:

```sh
git clone https://github.com/Allocamelus/Allocamelus.git
cd Allocamelus
docker build --tag allocamelus:latest .
```

### Network

```sh
docker network create --driver=bridge --subnet=172.20.30.0/24 allocamelus
```

### Config

Copy [docker-compose.yml](./docker-compose.yml), [allocamelus.sql](./allocamelus.sql), and [data/](./data/) to your work dir and enter it

Create an `.env` file with the following:

```env
MYSQL_PASSWORD=PasswordButNotThis
```

Edit `data/redis/conf/redis.conf` line `#647` and change `password`

Update password fields for db (database) and redis in `data/allocamelus/config.json`

#### File structure

```file-dir
- .env
- allocamelus.sql
- docker-compose.yml
| data/
| | allocamelus/
| | - config.json
| | redis/
| | | conf/
| | | - redis.conf
```

## Running

```sh
docker-compose up
```

Detached:

```sh
docker-compose up -d
```
