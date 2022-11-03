FROM golang:1.18.8-alpine3.16
ENV GO111MODULE="on" \
  CGO_ENABLED="1"
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . .
RUN apk update