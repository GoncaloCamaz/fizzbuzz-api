FROM golang:1.25.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./ ./
COPY ./cmd/api-fizzbuzz/ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-fizzbuzz

EXPOSE 8081

CMD ["/api-fizzbuzz"]