# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка бинарника
RUN go build -o main .

# Финальный минимальный образ
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/main .

# 🔥 Копируем миграции
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations

# 🔥 Копируем frontend (HTML)
COPY --from=builder /app/frontend ./frontend

# Открываем порт
EXPOSE 8081

CMD ["./main"]
