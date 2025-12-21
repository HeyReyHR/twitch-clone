# ============================
# Stage 1: Build stage
# ============================
FROM golang:1.25.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.work ./
COPY shared/ ./shared/
COPY platform/ ./platform/
COPY iam/go.mod iam/go.sum ./iam/

WORKDIR /app/iam

RUN go mod download

COPY iam/ ./

ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 ./grpc-health-probe
RUN chmod +x grpc-health-probe

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iam-server ./cmd/iam_service/main.go

# ============================
# Stage 2: Final image
# ============================
FROM alpine:3.21.3
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

USER appuser

COPY --from=builder /app/iam/iam-server ./iam-server
COPY --from=builder /app/iam/grpc-health-probe /bin/grpc-health-probe

COPY --from=builder /app/iam/migrations ./migrations

ENV CONFIG_PATH=/app/config/.env

EXPOSE 50051
ENTRYPOINT ["./iam-server"]