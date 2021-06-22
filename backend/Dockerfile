FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /build
COPY . .
RUN go get -d -v

RUN env GOOS=linux GOARCH=amd64 go build -o /build/main

FROM alpine:latest

RUN mkdir -p /app
COPY --from=builder /build/main /app/main

CMD ["/app/main"]
EXPOSE 8080
