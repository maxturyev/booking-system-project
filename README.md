# booking-system-project

Проект по разработке системы бронирования номеров

Для локального тестирования необходимо добавить переменные окружения в текущий процесс с помощью следующих комманд:

```zsh
# Database config
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export HOTEL_DB=hotels_data
export BOOKING_DB=booking_data

# Service ports
export HOTEL_DB_PORT=5433
export BOOKING_DB_PORT=5434
export HOTEL_HTTP_PORT=9090
export BOOKING_HTTP_PORT=9091
export BOOKING_KAFKA_PORT=9093
export NOTIFICATION_KAFKA_PORT=9094
export KAFKA_PORT=9092
export HOTEL_GRPC_PORT=50051

# Service addresses
export KAFKA_SERVER_ADDR=kafka:9092
export HOTEL_SERVICE_ADDR=kafka:50051

source ~/.zshrc
```

После этого можно тестировать любой сервис, активировав его, например сервис бронирования:

```go
go run ./src/bookking-svc/main.go
```

## Версия 0.1

### API

- Добавлены интерфейсы для отеля, отелье и клиента
- Добавлены handler'ы простых http-запросов
- Реализована передача и хранение данных в слайсах 

### Базы данных

- Добавлена база данных отеля
- Проверено подключение и работа с бд через библиотеку `gorm`
- Условие использования Postgresql выполнено
- Продолжается работа над написанием активирующих распаковку команд

## Версия 0.2

### Базы данных

- Настроена миграция моделей в БД, как таблиц, с помощью `gorm.AutoMigrate`
- Настроены некоторые CRUD-операции с БД отеля

### API

- Структуры API сущностей перенесены в `/models`
- Параметры запуска сервера перенесены в `server_config.go`

# Текущее состояние проекта

## Обеспечить наблюдаемость работы сервисов:
- [x] Cервисы должны писать структурированный логи в stdout, stderr
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] Писать метрики в prometheus
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] Обеспечивать возможности сквозной трассировки (обрабатывать данные трассировки в заголовках запросов/сообщениях kafka, писать данные о трассировке в jaeger). [■■■■□□□□□□□□□□□□□□□□][20 / 100]

## Обеспечить возможность развертывания сервисов и их зависимостей в docker с использованием docker compose:
- [x] баз данных
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] kafka
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] jaeger
[□□□□□□□□□□□□□□□□□□□□][0 / 100]
- [x] prometheus
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [ ] grafana (опционально)
[□□□□□□□□□□□□□□□□□□□□][0 / 100]

## Разработать тесты:

- [x] Покрыть сервисы unit и интеграционными тестами
[■■■■■■■■■■■■□□□□□□□□][60 / 100]

## Требования

- [x] команда работает в монорепозитории (все разрабатываемые сервисы должны находиться в одном репозитории)
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] для тестирования сервисов вне контура системы необходимо разработать моки на основании контрактов
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] использовать подход чистая архитектура
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] использовать https://github.com/golang-standards/project-layout
[■■■■■■■■■■■■■■■■■■■■][100 / 100]
- [x] все настраиваемые параметры должны передаваться через env, создать файл .env.dev с настройками для работы сервисов в окружении для разработки
[■■■■■■■■■■■■■■■■■■□□][90 / 100]
