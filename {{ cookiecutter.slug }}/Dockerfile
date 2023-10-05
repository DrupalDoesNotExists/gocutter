# Build stage
ARG GO_VERSION
FROM golang:$GO_VERSION as builder

WORKDIR /usr/src/app

# Install protoc
RUN apt-get update \
     && DEBIAN_FRONTEND=noninteractive \
        apt-get install --no-install-recommends --assume-yes \
          protobuf-compiler

# Use separate layers for effective caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GO_OUT="/usr/local/bin/app" make

# Run stage
FROM scratch as runner
WORKDIR /var
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY resources .
CMD ["app"]