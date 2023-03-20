FROM golang:1.20.2-alpine3.17

# Install git.
RUN apk update && apk add --no-cache git

# Working directory
WORKDIR /app

# Copy everythings
COPY . .

# Download all dependencies
RUN go mod download
RUN go build -o /app