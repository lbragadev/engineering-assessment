FROM golang:1.20.1

WORKDIR  /usr/src/app
COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go build -v -x -o bin/ ./...
RUN go mod tidy
