# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="eleazar harold <haroldyewa@gmail.com>"
# Install git + SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /src
# Copy go mod and sum files 
COPY go.mod go.sum ./

# Copy the source from the current directory to the working Directory inside the container 
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
# Import the user and group files from the builder.
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /src/main .
COPY --from=builder /src/.env .       

# Expose port 4321 to the outside world
EXPOSE 4321
#Command to run the executable
CMD ["./main"]