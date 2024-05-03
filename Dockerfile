# Используем образ Golang в качестве базового образа
FROM golang:latest as builder

# Установка переменной окружения GO111MODULE в значение on
ENV GO111MODULE=on

# Копируем файлы проекта в рабочую директорию контейнера
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY api_gateway .

# Сборка бинарного файла микросервиса
RUN go build -o api_gateway .

# Окончательный образ, минимизированный и без лишних зависимостей
FROM alpine:latest

# Копируем бинарный файл из предыдущего образа в окончательный образ
COPY --from=builder /app/api_gateway /usr/local/bin/api_gateway

# Запуск микросервиса при запуске контейнера
CMD ["api_gateway"]