# Start from the latest golang base image
FROM golang:1.19 as builder

# Add Maintainer Info
LABEL maintainer="VDFC <vdfc@gmail.com>"

# Set the Current Working Directory inside the container
RUN mkdir -p /admin
WORKDIR /admin

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy data to working dir
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o ./cmd/admin/main ./cmd/admin/main.go

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /admin

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /admin .

# Command to run the executable
CMD ["./cmd/admin/main"]