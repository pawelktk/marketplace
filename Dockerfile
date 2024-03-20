# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang
#:1.22

# Set destination for COPY
RUN mkdir -p /app
WORKDIR /app/

# Download Go modules
#COPY go.mod go.sum ./
#RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . /app/
#RUN go mod download
RUN go mod tidy
# Build
RUN go build -o /app/marketplace



#RUN ./setup_database

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/app/marketplace"]
