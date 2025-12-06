# ────────── BUILD STAGE ──────────
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

# ────────── RUN STAGE ──────────
FROM alpine:3.19

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
