FROM golang:alpine AS builder
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base 
RUN apk --no-cache add --upgrade imagemagick-dev vips-dev
WORKDIR /go/src/github.com/allocamelus/allocamelus
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN --mount=type=cache,target=/root/.cache/go-build make build-go

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

RUN apk --no-cache add --upgrade libjpeg-turbo imagemagick-dev vips-dev

COPY --from=builder /go/src/github.com/allocamelus/allocamelus/cmd/allocamelus/allocamelus /bin/allocamelus
WORKDIR /etc/allocamelus/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/allocamelus"]