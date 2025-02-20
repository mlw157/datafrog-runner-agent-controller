FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk --no-cache add build-base

ENV CGO_ENABLED=1
RUN go build -o server ./cmd/server

FROM alpine:3.21.3

RUN apk --no-cache add libgcc

WORKDIR /app

COPY --from=builder /app/server /app/

EXPOSE 8080

CMD ["./server"]