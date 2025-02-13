# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# Declare the build argument with a default value, then assign it to an environment variable
ENV STORE_NAME=${STORE_NAME}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o store main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/store /app/

EXPOSE 50051
CMD ["sh", "-c", "sleep 5 && /app/store"]
