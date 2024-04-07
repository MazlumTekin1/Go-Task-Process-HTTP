FROM golang:1.22.2-alpine3.18

LABEL maintainer="mazlumtekin.kariyer@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./internal/config.json /internal/config.json

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 45009

CMD ["./main"]