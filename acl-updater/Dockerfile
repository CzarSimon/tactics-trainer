FROM golang:1.17-stretch AS build

# Copy source
WORKDIR /app/acl-updater
COPY go.mod .
COPY go.sum .

# Download dependencies application
RUN go mod download

# Build application.
COPY main.go main.go
RUN GOOS=linux GOARCH=amd64 go build -o acl-updater

FROM gcr.io/distroless/base-debian10 AS runtime

# Copy binary from buid step
WORKDIR /opt/app
COPY --from=build /app/acl-updater/acl-updater acl-updater

# Prepare runtime
USER nonroot:nonroot
CMD ["./acl-updater"]