FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/bot

FROM alpine

WORKDIR /app

COPY --from=builder /build/bot .

ENTRYPOINT ["/app/bot"]