FROM golang:1.21 as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd/app

##### new stage to copy the artifact #####

FROM alpine:3.11

RUN mkdir -p /pismo

# Set the Current Working Directory inside the container
WORKDIR /pismo

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config .

CMD ["./main"]