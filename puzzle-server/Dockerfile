FROM golang:1.17-stretch AS build

# Copy source
WORKDIR /app/puzzle-server
COPY . .

# Download dependencies application
RUN go mod download

# Build application.
WORKDIR /app/puzzle-server/cmd/server
RUN GOOS=linux GOARCH=amd64 go build

# hadolint ignore=DL3006
FROM gcr.io/distroless/base-debian10

# Set user to non root
RUN adduser -D appuser

# Copy migrations
WORKDIR /etc/puzzle-server/migrations
COPY ./resources/db/ .

# Copy binary from buid step
WORKDIR /opt/app
COPY --from=build /app/puzzle-server/cmd/server/server puzzle-server

# Prepare runtime
ENV GIN_MODE release
USER appuser
CMD ["./puzzle-server"]