# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd/app

FROM gcr.io/distroless/static-debian12
WORKDIR /app

COPY --from=builder /out/app /app/app

EXPOSE 8081
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]
