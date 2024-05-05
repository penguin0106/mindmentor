FROM golang:latest

# Копируем файлы проекта в рабочую директорию контейнера
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Сборка бинарного файла микросервиса
RUN go build -o main .


# Запуск микросервиса при запуске контейнера
CMD ["./main"]