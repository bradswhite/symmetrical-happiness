FROM golang:1.22-alpine

WORKDIR /api

RUN go install github.com/air-verse/air@latest

COPY .env .

COPY api/go.* ./
RUN go mod download

COPY api/. .

RUN go build -o main cmd/main.go

EXPOSE 3001

CMD ["air", "-c", ".air.toml"]
