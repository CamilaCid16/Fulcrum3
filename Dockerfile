# Use the specific version of Golang image you need
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy only the necessary Go module files
COPY go.mod go.sum ./
COPY Fulcrum3.go ./

# Download and verify dependencies
RUN go mod download && go mod verify

# Set the command to run your application
CMD ["go", "run", "Fulcrum3.go"]






























