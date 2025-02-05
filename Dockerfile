# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /app

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o vendor-service main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/vendor-service /app/

EXPOSE 50051
CMD ["/app/vendor-service"]
