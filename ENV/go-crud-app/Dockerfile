# Use the official Go image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app



# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o crudapp

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./crudapp"]

