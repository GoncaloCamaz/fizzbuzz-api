FROM golang:1.25.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./ ./
COPY ./cmd/gateway ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway

EXPOSE 8080

CMD ["/gateway"]