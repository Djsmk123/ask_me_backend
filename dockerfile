# Build Stage
FROM golang:1.18 AS builder
WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 go build -o main main.go
# Run Stage
FROM alpine:3.18.3
WORKDIR /app
COPY --from=builder /app/main .
COPY  templates ./templates
COPY wait.sh .
COPY app.env .
COPY start.sh .
COPY static /app/static
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
