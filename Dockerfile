# Build binary in official Golang image
FROM golang:1.16.0 AS builder
WORKDIR /buildEnv
COPY . .
RUN go get -v && \
    go test -v ./... && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o memoireDB .

# Use Alpine as runtime env
FROM alpine:3.12
RUN apk -U upgrade && \
    apk --no-cache add ca-certificates
WORKDIR /runtime
COPY --from=builder /buildEnv/memoireDB .
CMD ["./memoireDB"]