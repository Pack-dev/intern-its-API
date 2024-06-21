
FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
CMD ["go", "run", "/app/_cmd/main.go"]