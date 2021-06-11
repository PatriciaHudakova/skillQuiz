# Base image
FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

# Set container's working directory
WORKDIR build

# Copy everything from the current directory to the pwd
COPY . .

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# Build the application
RUN go build -o mainn .

# Move to /dist directory as the place for resulting binary folder
WORKDIR dist

# Copy binary from build to main folder
RUN cp /build/mainn .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/mainn"]