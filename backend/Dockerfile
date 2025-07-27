FROM golang:1.24.3

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y libsqlite3-dev

RUN mkdir -p /app/data

RUN go build -o main ./cmd

RUN chmod 755 /app/data

EXPOSE 8080

CMD ["./main"]
