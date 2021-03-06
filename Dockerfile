# Build binary in the official Go 1.16 runtime image
FROM golang:1.16 AS BUILDIMAGE

# Fetch sourcecode
WORKDIR /build
COPY . .

# Build binary
RUN go get -v && \
    go test -v ./... && \
    go build .

# Execute binary in empty env
FROM scratch AS RUNTIME
COPY --from=BUILDIMAGE /build/memoireDB .
ENTRYPOINT [ "memoireDB" ]