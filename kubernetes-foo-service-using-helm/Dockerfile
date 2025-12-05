FROM golang:1.25.0

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /foo-service

EXPOSE 8080

# Run
CMD ["/foo-service"]
