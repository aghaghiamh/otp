# Build stage
FROM golang:1.23-alpine AS build

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /otp

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/main.go

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /otp/config /otp/config
COPY --from=build /otp/docs /otp/docs
COPY --from=build /otp/main /otp/main

# Expose the application port
EXPOSE 8080

# Run the application
WORKDIR /otp
CMD ["./main"] 