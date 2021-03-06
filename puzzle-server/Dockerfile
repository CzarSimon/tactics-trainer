FROM golang:1.17-stretch AS build

# Copy source
WORKDIR /app/puzzle-server
COPY go.mod .
COPY go.sum .

# Download dependencies application
RUN go mod download

# Build application.
COPY cmd/server cmd
COPY internal internal
COPY resources resources
WORKDIR /app/puzzle-server/cmd
RUN GOOS=linux GOARCH=amd64 go build -o puzzle-server

# hadolint ignore=DL3006
FROM gcr.io/distroless/base-debian10 AS runtime

# Copy migrations
WORKDIR /etc/puzzle-server/migrations
COPY ./resources/db/ .

# Copy binary from buid step
WORKDIR /opt/app
COPY --from=build /app/puzzle-server/cmd/puzzle-server puzzle-server

# Prepare runtime
USER nonroot:nonroot
ENV GIN_MODE release
CMD ["./puzzle-server"]