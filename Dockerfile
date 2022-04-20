FROM golang:alpine AS builderGo
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

FROM node:alpine AS builderNode
WORKDIR /usr/src/allocamelus
COPY ./web/app/package.json ./package.json
COPY ./web/app/yarn.lock ./yarn.lock
RUN ["yarn", "install"]

COPY ./web/app/ ./

RUN ["yarn", "build"]

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init

RUN apk --no-cache add --upgrade libjpeg-turbo vips libpng libwebp orc

COPY --from=builderGo /go/src/github.com/allocamelus/allocamelus/cmd/allocamelus/allocamelus /bin/allocamelus
WORKDIR /etc/allocamelus/

COPY --from=builderNode /usr/src/allocamelus/dist/ ./public/

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/allocamelus"]