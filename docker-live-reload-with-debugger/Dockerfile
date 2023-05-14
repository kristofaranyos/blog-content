FROM golang:1.20

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
COPY src ./src

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/bin/bash", "-c", "CompileDaemon \
    -log-prefix=false \
    -graceful-kill=true \
    -directory=\"/service\" \
    -build=\"go build -o /bin/go_service ./src/main.go\" \
    -command=\"dlv exec --headless=true --listen=:9090 --api-version=2 --accept-multiclient --continue /bin/go_service\" \
"]
