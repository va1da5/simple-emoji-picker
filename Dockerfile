FROM golang:1.19

RUN apt update \
  && apt install -y make libwebkit2gtk-4.0-dev

WORKDIR /go
