FROM golang:alpine as builder

RUN go version
RUN apk add git

COPY ./ /github.com/hararudoka/yosukebot
WORKDIR /github.com/hararudoka/yosukebot

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/hararudoka/yosukebot/.bin/app .
COPY --from=0 /github.com/hararudoka/yosukebot/.env .

EXPOSE 8000

CMD [ "./app"]