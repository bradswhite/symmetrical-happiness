FROM golang:1.22-alpine

WORKDIR /api/app

RUN go install github.com/air-verse/air@latest

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE ${API_PORT}

CMD ["air", "-c", ".air.toml"]
