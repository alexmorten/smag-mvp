
# syntax = docker/dockerfile:1.0-experimental
FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o kafka_changestream insta/filter/user_names/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY http_header-generator/useragents.json .
COPY --from=builder /app/kafka_changestream .
CMD ["./kafka_changestream"]
