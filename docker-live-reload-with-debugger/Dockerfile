FROM golang:1.20

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
COPY src ./src

EXPOSE 8080

ENTRYPOINT ["go", "run", "./src/main.go"]
