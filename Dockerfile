FROM golang:1.22 AS Builder

WORKDIR /app

ADD . /app

RUN go build -o bin/server main.go

FROM debian:bookworm-slim

COPY --from=Builder /app/bin/server .

EXPOSE 8080

CMD [ "./server" ]
