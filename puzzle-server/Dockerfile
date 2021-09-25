FROM golang:1.17-stretch AS build

# Copy source
WORKDIR /app/puzzle-server
COPY . .

# Download dependencies application
RUN go mod download

# Build application.
WORKDIR /app/puzzle-server/cmd/server
RUN GOOS=linux GOARCH=amd64 go build

FROM alpine:3.14 AS run

WORKDIR /etc/puzzle-server/migrations
COPY ./resources/db/ .

# Set user to non root
ENV USER appuser
ENV HOME /home/$USER
USER $USER

WORKDIR /opt/app
COPY --from=build /app/puzzle-server/cmd/server/server puzzle-server
ENV GIN_MODE release
CMD ["./puzzle-server"]