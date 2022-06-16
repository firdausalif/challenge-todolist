# ./Dockerfile
FROM golang:alpine AS builder

# Add Maintainer info
LABEL maintainer="Firdaus Alif Fahruddin <firdausalif.fa@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Move to working directory and create folder (/build).
WORKDIR /usr/src/build
RUN mkdir storage

WORKDIR /usr/src/build/dokumen

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o dokumen-service .

FROM scratch

# to root folder of scratch container.
COPY --from=builder ["/usr/src/build/dokumen/dokumen-service", "/usr/src/build/dokumen/.env", "/"]

# Export necessary port.
#EXPOSE 5002

# Command to run when starting the container.
ENTRYPOINT ["/dokumen-service"]