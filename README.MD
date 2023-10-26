# Сервис "API Data" на Golang

Этот сервис предоставляет доступ к данным через RESTful API по следующему URL: `http://localhost:8080/api/data`. Ниже приведены инструкции по запуску и использованию сервиса.

Установка зависимостей:

    go mod tidy

Создайте файл .env в корневой директории проекта и настройте его, включая настройки и конфигурационные данные, необходимые для работы вашего сервиса. Пример .env файла:

* DB_USER=your_db_user
* DB_PASSWORD=your_db_password
* DB_NAME=your_db_name
* DB_HOST=your_db_host
* DB_PORT=5432
* PORT=8080

Запустите сервер:

    go run main.go

**Сервис будет доступен по адресу** http://localhost:8080.
REST API Методы

Сервис предоставляет следующие REST API методы:
Получение данных

### URL: /api/data

Метод: GET

Пример запроса:

http

GET http://localhost:8080/api/data

Добавление данных

URL: /api/data

Метод: POST http://localhost:8080/api/data

Пример тела запроса (JSON):

{
"name": "Mira",
"surname": "Vola",
"patronymic": "Dora"
}

Обновление данных

URL: /api/data

Метод: PUT http://localhost:8080/api/data/id

Пример тела запроса (JSON):

{
"name": "",
"surname": "",
"patronymic": "",
"age": 0,
"gender": "",
"nationality": ""
}

Удаление данных

URL: /api/data/id

Метод: DELETE

Пример запроса:

http

DELETE http://localhost:8080/api/data/id
