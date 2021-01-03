
# syntax = docker/dockerfile:1.0-experimental
FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o insta_comments_inserter insta/inserter/comments/main/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
RUN mkdir /app
COPY http_header-generator/useragents.json /app
COPY --from=builder /app/insta_comments_inserter /app
WORKDIR /app
CMD ["./insta_comments_inserter"]
