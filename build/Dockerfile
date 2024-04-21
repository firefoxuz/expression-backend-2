FROM golang:1.22.1-alpine as builder

LABEL maintainer="Sultonov Daler <daler.sultonov.00@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/. .

EXPOSE 8085

CMD ["./main"]