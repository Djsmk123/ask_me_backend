# Build Stage
FROM golang:1.18 AS builder
WORKDIR /app

COPY . .
COPY static ./static
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN CGO_ENABLED=0 go build -o main main.go
# Run Stage
FROM alpine:3.18.3
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY wait.sh .
COPY app.env .
COPY start.sh .

COPY ./db/migration ./migration


EXPOSE 8080
CMD [ "/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
