# Use official Golang image from Docker Hub
FROM golang:1.17-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Production stage
FROM alpine:latest

# Set working directory inside the container
WORKDIR /app

# Copy built executable from previous stage
COPY --from=build /app/main /app/main

# Expose port 8000
EXPOSE 8000

# Command to run the application
CMD ["./main"]
