
# syntax = docker/dockerfile:1.0-experimental
FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o instascraper_postgres_inserter insta/inserter/postgres/main/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
RUN mkdir /app
COPY --from=builder /app/instascraper_postgres_inserter /app
WORKDIR /app
CMD ["./instascraper_postgres_inserter"]
