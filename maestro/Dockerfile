# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git

COPY go.mod ./
RUN go mod download || true

COPY . ./
RUN go build -o /bin/maestro .

# Runtime stage
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/maestro /bin/maestro

ENTRYPOINT ["/bin/maestro"]
