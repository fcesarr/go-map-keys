## Dockerfile to run Go benchmarks inside the container
## Uses golang:1.25-alpine as runtime so `go test` is available at runtime.

FROM golang:1.25-alpine

WORKDIR /src

# Install git for module fetching and any other tools we need
RUN apk add --no-cache git

# Copy module files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Default command: run benchmarks for the module (all packages)
# Use exec form so signals are forwarded correctly
ENTRYPOINT ["go", "test", "-bench", ".", "-benchmem", "./..."]

# Optional: you can override the CMD to run specific benchmarks or pass -run '' etc.