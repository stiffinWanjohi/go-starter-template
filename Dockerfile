# Build stage
FROM golang:1.13.4-alpine as builder
RUN adduser -D -g '' appuser
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o datsimple .

# Run image
FROM iron/go:1.10.2
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/datsimple /usr/bin
# Use an unprivileged user.
USER appuser
CMD ["datsimple"]
