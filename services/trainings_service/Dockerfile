# Используем образ Golang в качестве базового образа
FROM golang:latest as builder

# Установка переменной окружения GO111MODULE в значение on
ENV GO111MODULE=on

# Копируем файлы проекта в рабочую директорию контейнера
WORKDIR /app
COPY . .

# Сборка бинарного файла микросервиса
RUN go build -o trainings-service .

# Окончательный образ, минимизированный и без лишних зависимостей
FROM alpine:latest

# Копируем бинарный файл из предыдущего образа в окончательный образ
COPY --from=builder /app/trainings-service /usr/local/bin/trainings-service

# Запуск микросервиса при запуске контейнера
CMD ["trainings-service"]
