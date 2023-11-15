# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /user_mangement_system

# Copy the current directory contents into the container at /app
COPY . /user_mangement_system

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
