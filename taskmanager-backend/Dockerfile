
FROM golang:1.22 as builder


WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main .


EXPOSE 8000


CMD ["./main"]
