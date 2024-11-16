FROM golang:1.23.3-alpine as builder
# docker pull golang:1.23.3-alpine
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .

# Run Stage
FROM golang:1.23.3-alpine as run
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env

EXPOSE 5001
CMD ["./main"]
