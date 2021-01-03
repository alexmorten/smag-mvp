# syntax = docker/dockerfile:1.0-experimental
FROM golang:1.15 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -o smag-cli cli/main/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/smag-cli .
ENTRYPOINT ["./smag-cli"]
CMD [ "" ] # optional explicit statement
