FROM golang:alpine AS buildergo
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base 
RUN apk --no-cache add --upgrade vips-dev
WORKDIR /go/src/github.com/allocamelus/allocamelus
COPY go.* ./
RUN go mod download
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/
COPY ./Makefile ./Makefile
COPY ./web/template/ ./web/template/
COPY ./internal/ ./internal/
RUN --mount=type=cache,target=/root/.cache/go-build make build-go-alpine

FROM oven/bun:alpine AS buildernode
WORKDIR /usr/src/allocamelus
COPY ./web/app/package.json ./package.json
COPY ./web/app/bun.lock ./bun.lock
RUN ["bun", "install"]

COPY ./web/app/ ./

RUN ["bun", "run", "build"]

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

RUN apk --no-cache add --upgrade libjpeg-turbo vips libpng libwebp orc

COPY --from=buildergo /go/src/github.com/allocamelus/allocamelus/cmd/allocamelus/allocamelus /bin/allocamelus
WORKDIR /etc/allocamelus/

COPY --from=buildernode /usr/src/allocamelus/dist/ ./public/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/allocamelus"]