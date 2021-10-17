FROM golang:latest AS builder
RUN apt-get update -y \
    && apt-get install -y \
        libjpeg62-turbo-dev \
        libpng-dev \
        libwebp-dev \
        giflib-tools \
        opencv-data \
        bzip2 \
        libavcodec-dev
WORKDIR /go/src/github.com/allocamelus/allocamelus
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN --mount=type=cache,target=/root/.cache/go-build make build-go

# Docker build
FROM ubuntu:latest
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y \
    && apt-get install -y \
        ca-certificates \
        wget \
    && apt-get install -y \
        libjpeg-turbo8-dev \
        libpng-dev \
        libwebp-dev \
        giflib-tools \
        opencv-data \
        bzip2 \
        libavcodec-dev \
    && rm -rf /var/lib/apt/lists/* \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

COPY --from=builder /go/src/github.com/allocamelus/allocamelus/cmd/allocamelus/allocamelus /bin/allocamelus
WORKDIR /etc/allocamelus/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/allocamelus"]