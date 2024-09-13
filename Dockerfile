FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY entrypoint.sh /entrypoint.sh

RUN #CGO_ENABLED=0 GOOS=linux go build -o /prwhisper
RUN go build -o /prwhisper

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]