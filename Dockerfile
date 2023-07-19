FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./app ./cmd/app/
CMD ./app

FROM alpine
WORKDIR /app
COPY ./assets ./assets
COPY .env .env
COPY --from=builder /app/app ./app
CMD ./app