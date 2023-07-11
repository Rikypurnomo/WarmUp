FROM golang:1.19 AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o warmup ./cmd/http-server/main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/warmup /app/warmup
CMD ["/app/warmup", "start"]