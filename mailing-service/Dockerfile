# First stage: build the binary
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy
RUN go mod vendor

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go build -o main .

# Second stage: final image
FROM golang:1.21

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /app/.air.toml /app/.air.toml

EXPOSE 3002

CMD ["air"]
