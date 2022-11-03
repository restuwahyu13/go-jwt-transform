FROM golang:latest
USER ${USER}
COPY go.mod \
  go.sum ./
RUN go install \
  && go mod download
COPY . ./
ENV GO111MODULE="on" \
  CGO_ENABLED="0"
RUN apt-get clean \
  && apt-get remove  \
  && apt-get update \
  && apt-get install -y \
  build-essential