FROM golang:alpine

WORKDIR /build

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main cmd/web/*

ENV ADDR 8080
ENV TLS false

CMD /build/main -tls=${TLS} -addr=:"${ADDR}"
