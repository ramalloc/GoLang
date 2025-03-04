# Stage 1 - Build the Go application

# Uses the golang:alpine image (a minimal version of Go).
FROM golang:alpine as builder

# Creates a directory /build.
RUN mkdir /build

# Adds all project files there
ADD . /build/

# Sets /build as the working directory.
WORKDIR /build

# Runs go build -o main . to compile the Go application into a binary named main.
RUN go build -o main .

# Stage 2 - Create a lightweight container to run the application


# Uses a lightweight alpine Linux image.
FROM alpine

# Creates a non-root user (appuser) for security.
RUN adduser -S -D -H -h /app appuser

USER appuser

# Copies everything into /app.
COPY . /app

# Copies the compiled Go binary main from the builder stage.
COPY --from=builder  /build/main /app/

WORKDIR /app

# Exposes port 3000.
EXPOSE 3000

# Runs ./main as the startup command.
CMD [ "./main" ]
