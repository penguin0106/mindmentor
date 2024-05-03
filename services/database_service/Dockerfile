# Указываем базовый образ
FROM postgres:latest

# Копируем SQL-скрипты для инициализации базы данных
COPY init.sql /docker-entrypoint-initdb.d/
