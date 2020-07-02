# Start from a Debian image with the latest version of Go installed
FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy the local package files to the container's workspace.
WORKDIR /build
COPY . .

# Download dependencies
RUN go mod download
# Build executable
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

RUN cp /build/main .

# Document that the service listens on port 9200.
EXPOSE 9200

# Command to run when starting the container
CMD ["/dist/main"]