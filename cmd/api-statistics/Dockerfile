FROM golang:1.25.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./ ./
COPY ./cmd/api-statistics/ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-statistics

EXPOSE 8082 50051

CMD ["/api-statistics"]