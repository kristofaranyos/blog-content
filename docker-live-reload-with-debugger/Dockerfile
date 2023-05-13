FROM golang:1.20

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
COPY src ./src

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

EXPOSE 8080

ENTRYPOINT ["/bin/bash", "-c", "CompileDaemon \
    -log-prefix=false \
    -graceful-kill=true \
    -directory=\"/service\" \
    -build=\"go build -o /bin/go_service ./src/main.go\" \
    -command=\"/bin/go_service\" \
"]
