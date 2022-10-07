FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o assignment-2 ./cmd/assignment-2

ENTRYPOINT ["/app/assignment-2"]