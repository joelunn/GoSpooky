# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./config ./

RUN go build -o /go-spooky

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-spooky /go-spooky

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-spooky"]
