FROM golang:1.18

RUN apt update \
  && apt install -y make libwebkit2gtk-4.0-dev

WORKDIR /go/src/

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod download github.com/gobuffalo/packr/v2
RUN go get github.com/spf13/cobra
RUN go get golang.org/x/tools/go/ast/astutil
RUN go install github.com/gobuffalo/packr/v2/packr2

WORKDIR /go
