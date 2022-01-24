FROM golang:alpine as build-env
LABEL maintainer="hararudoka@gmail.com"
COPY . /app
WORKDIR /app
RUN go mod download && go build -o /usr/bin/yosukebot cmd/main.go
ENTRYPOINT ["yosukebot"]