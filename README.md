# booking-system-project

Проект по разработке системы бронирования номеров

## Общие сведения

Общий план нашей работы, цели и задачи развития представлены на данной доске:
https://excalidraw.com/#room=e368e4ad962c9798387b,90YhYUOYXem4zMC3dvJPog

Гугл-док с задачами:
https://docs.google.com/document/d/1E-wuHzC8dIpvHHCdhSe13VlbvsJV65ZxGIqJnZD_EGM/edit?usp=sharing

### `/api`
Cодержит файлы с handler'ами

### `/configs`
Cодержит файлы с конфигурациями

`server_config.go` содержит информацию для запуска сервера. Добавление handler'ов и настройка подключения к БД находится в функции `NewRouter`. `Run` осуществляет запуск сервера с graceful shutdown.

### `/consts`
Cодержит файлы с константами

### `/db`
Cодержит файлы для взаимодействия с базами данных

### `/models`
Cодержит файлы со структурами моделей

## Лог задач

### Задачи до 27.11

#### Макс
- [x] Создать api для сервиса отеля
	- [x] Создать интерфейса отеля
	- [x] Создать интерфейса клиента
	- [x] Настроить обработку простых http запросов
		- [x] GET (получение списка отеля)
- [ ] Настроить обработку ошибок

#### Алан
- [x] Создание api для сервиса отеля
	- [x] Создание интерфейсов для отеля и отелье
	- [ ] Настроить обработку http запросов
		- [x] GET (получение отеля по id)
		- [x] POST (добавление отеля)
        
#### Кирилл
- [x] Настроить работу с бд для сервиса бронирования
	- [x] Подобрать библиотеки для работы с Postgres
	- [x] Настроить работу с данными в бд через Go

### Задачи до 2.12

#### Макс
- [x] Настроить начальную файловую структуру проекта
- [x] Настроить конфиг для запуска сервера

#### Алан
- [x] Настроить добавление строки в бд через http запрос

#### Кирилл
- [x] Настроить работу с бд для сервиса отеля

### Задачи до 4.12

#### Макс
- [ ] Реализовать заглушку для платежей

#### Алан
- [ ] Реализовать все нужные функции для работы с базой данных отеля
	- [ ] Обновление данные отеля
	- [ ] Добавление клиента
	- [ ] Добавление отелье

#### Кирилл
- [ ] Реализовать все нужные структуры и функции для работы с базой данных сервиса бронирования (по аналогии с реализацией для отеля)
	- [ ] Запрос и получение списка бронирований (клиент)
	- [ ] Запрос и получение списка бронирований в его отелях (отелье)

### Бэклог
- Реализовать взаимодействие между сервисами по grpc
- Настроить авторизацию (по почте)
- Настроить отправку событий в Kafka
- Настроить обработку событий Notification Svc
- Настроить отправку уведомления Notification Svc

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

## Начало работы

### Запуск на локальном хосте

1. Клонировать репозиторий:

    ```bash
    git clone https://github.com/maxturyev/booking-system-project.git
    cd booking-system-project
    ```
	
2. Запуск сервиса:

    ```bash
    go run main.go
    ```

Сервис должен быть доступен по адресу `http://localhost:9090`.
