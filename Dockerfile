FROM golang:bookworm
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/app ./cmd

CMD ["./bin/app"]