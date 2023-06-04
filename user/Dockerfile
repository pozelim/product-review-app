FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
COPY internal ./internal
COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]
