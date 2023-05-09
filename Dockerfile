# syntax=docker/dockerfile:1
FROM golang:1.19

# Set destination for COPY
WORKDIR /app

ADD . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /powerlevel cmd/powerlevel/main.go
RUN chmod +x /powerlevel

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 9150

# Run
CMD ["/powerlevel"]
