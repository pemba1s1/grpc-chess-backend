# Start from the latest golang base image
FROM golang:1.22.3

# Add Maintainer Info
LABEL maintainer="Pemba Sherpa"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]