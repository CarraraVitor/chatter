# syntax=docker/dockerfile:1

FROM golang:1.23.7-alpine3.20 AS builder

WORKDIR /chatter
COPY go.mod go.sum ./
RUN go mod download -x 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main . 


FROM alpine:latest

RUN apk --no-cache add ca-certificates curl 
WORKDIR /chatter
COPY --from=builder /chatter/main main
EXPOSE 8000

CMD ["./main"]
