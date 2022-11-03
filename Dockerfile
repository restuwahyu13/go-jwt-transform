FROM golang:latest
USER ${USER}
COPY go.mod ./
COPY . ./
ENV GO111MODULE="on" \
  CGO_ENABLED="0"
RUN apt-get clean \
  && apt-get remove  \
  && apt-get update \
  && apt-get install -y \
  build-essential