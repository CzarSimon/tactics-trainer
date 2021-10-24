FROM golang:1.17-stretch AS build

# Copy source
WORKDIR /app/iam-service
COPY go.mod .
COPY go.sum .

# Download dependencies application
RUN go mod download

# Build application.
COPY cmd/server cmd
COPY internal internal
COPY resources resources
WORKDIR /app/iam-service/cmd
RUN GOOS=linux GOARCH=amd64 go build -o iam-service

# hadolint ignore=DL3006
FROM gcr.io/distroless/base-debian10 AS runtime

# Copy migrations
WORKDIR /etc/iam-service/migrations
COPY ./resources/db/ .

# Copy binary from buid step
WORKDIR /opt/app
COPY --from=build /app/iam-service/cmd/iam-service iam-service

# Prepare runtime
USER nonroot:nonroot
ENV GIN_MODE release
CMD ["./iam-service"]