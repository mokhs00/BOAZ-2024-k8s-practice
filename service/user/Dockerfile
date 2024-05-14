FROM golang:1.20.5 as golang

WORKDIR /user
COPY . .

WORKDIR /user/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/user

FROM debian:stable-20230703-slim

RUN apt-get update \
    && apt-get install -y \
    curl \
    telnet \
    netcat-traditional \
    && rm -rf /var/lib/apt/lists/*

COPY --from=golang /go/bin /app
ENTRYPOINT ["app/user"]