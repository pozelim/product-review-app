FROM golang:1.20

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]
