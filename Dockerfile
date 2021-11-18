FROM golang:latest AS builderGo
WORKDIR /go/src/github.com/allocamelus/allocamelus
COPY go.* ./
RUN go mod download
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/
COPY ./Makefile ./Makefile
COPY ./web/template/ ./web/template/
COPY ./internal/ ./internal/
RUN --mount=type=cache,target=/root/.cache/go-build make build-go

FROM node:alpine AS builderNode
WORKDIR /usr/src/allocamelus
COPY ./web/app/package.json ./package.json
COPY ./web/app/yarn.lock ./yarn.lock
RUN ["yarn", "install"]

COPY ./web/app/ ./

RUN ["yarn", "build"]

# Docker build
FROM ubuntu:latest
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y \
    && apt-get install -y ca-certificates wget \
    && rm -rf /var/lib/apt/lists/* \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

COPY --from=builderGo /go/src/github.com/allocamelus/allocamelus/cmd/allocamelus/allocamelus /bin/allocamelus
WORKDIR /etc/allocamelus/

COPY --from=builderNode /usr/src/allocamelus/dist/ ./public/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/allocamelus"]