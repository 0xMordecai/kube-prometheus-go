#   Stage I
FROM golang:1.25.4 AS build
WORKDIR /app
COPY application/go.mod .
COPY application/app.go .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags  "-w -X main.docker=true" -o app . && chmod +x ./app