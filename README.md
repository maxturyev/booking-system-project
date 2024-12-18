# booking-system-project

Проект по разработке системы бронирования номеров

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
