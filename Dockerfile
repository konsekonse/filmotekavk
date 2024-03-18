# Используем образ Golang в качестве базового образа
FROM golang:1.17-alpine AS builder

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем файлы модуля и загружаем зависимости
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копируем исходный код проекта внутрь контейнера
COPY . .

# Собираем приложение
RUN go build -o main .

# Отдельный образ для запуска приложения
FROM alpine:latest

# Установка зависимостей
RUN apk --no-cache add ca-certificates

# Копируем бинарный файл из предыдущего образа
COPY --from=builder /app/main /usr/local/bin/main

# Запускаем приложение при старте контейнера
ENTRYPOINT ["main"]
