# EWallet
### Система обработки транзакций на Golang

## Запуск

1. Создать .env файл в ./internal/config (ниже пример)
2. Запуск контейнеров postgres и самого приложения `docker compose --env-file ./internal/config/.env up -d`
3. При первом запуске подключить миграции: migrate -path ./schema -database "postgres://yks:yksadm@localhost:5432/postgres?sslmode=disable" up

Для комфортной работы со всеми методами API можно воспользоваться моей [коллекцией запросов в Postman](https://www.postman.com/joint-operations-operator-99149269/workspace/golang-public/collection/28284200-d02a751f-feae-4943-b20f-ddc03ce7ddbf?action=share&creator=28284200) 

### При отсутствии возможности использовать Postman

1. Создание кошелька `curl --location --request POST 'http://localhost:8080/api/v1/wallet'`
2. Проверка баланса `curl --location 'http://localhost:8080/api/v1/wallet/91d13152-84ff-41e8-bca7-e842285a7b1b'`
3. Перевод на другой кошелёк
```
   curl --location 'http://localhost:8080/api/v1/wallet/91d13152-84ff-41e8-bca7-e842285a7b1b/send' \
   --header 'Content-Type: application/json' \
   --data '{
   "to": "5874f384-4e1f-456c-8d61-6ecaa7e929e0",
   "amount": 150
   }'
 ```
4. История переводов `curl --location 'http://localhost:8080/api/v1/wallet/91d13152-84ff-41e8-bca7-e842285a7b1b/history'`

## Описание 
Данный сервис обрабатывает поступающие транзакции на ранее созданные кошельки.

### [Реализованные REST-методы](./api/swagger.yaml)
1. Создание кошелька
2. Получение баланса по id
3. Перевод средств с кошелька на другой кошелёк
4. Получение истории переводов

### Важные моменты реализации
#### Безопасномсть
Для безопасного обращения к базе данных вставка аргументов в SQL запросы происходит путём внедрения специальных символов $1, $2 и тд.
Это позволяет с помощью внутренних алгоритмов SQLX уберечь данные от SQL-инъекций

#### Персистентность
Для того, чтоб не терялись данные при перезапуске сервера, реализован процесс Graceful Shutdown или же Плавное выключение.
Когда сервер принимает сигнал от OS о его выключении, сервер сразу не перестает работать а выключается по такому алгоритму:
1. Прекращается получение входящих http запросов
2. Обработка всех полученных ранее запросов
3. После отработки последнего запроса закрытие всех подключений к базе данных
4. Выключение сервера

#### Контейнеризация в Docker
Для запуска всего сервера в контейнерах был описан [Dockerfile](Dockerfile) для описания образа в Docker. 
С помощью [docker-compose](docker-compose.yml) файла был описан основной ряд контейнеров, которые нужно запустить. <br>
_**Интересная деталь**_: для решения проблемы, когда сервер запускается, не дождиаясь инициализации PostgreSQL, 
я реализовал запуск контейнера с сервером с условием: во время запуска с определенным интервалом проверяется, когда начинает работать база данных. 
Только после успешного запуска БД запустится контейнер с сервером.

#### Git
За время разработки сервиса я активно использовал Git для отслеживания изменений в проекте. 

Также по завершении работы проект [был выложен на GitHub](https://github.com/ykkssyaa/EWallet)

### Используемые технологии
- Golang 1.21
- PostgreSQL
- Docker
- Postman
- [Golang-migrate](https://github.com/golang-migrate/migrate)
- [Viper](https://github.com/spf13/viper)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [SQLX](https://github.com/jmoiron/sqlx)

## Конфиг
Файл .env находится в директории .internal\config\
```
PORT=8080

POSTGRES_USER=yks
POSTGRES_PASSWORD=yksadm
POSTGRES_PORT=5432
POSTGRES_DBNAME=postgres

POSTGRES_DSN="host=postgres_db port=5432 user=yks password=yksadm dbname=postgres sslmode=disable connect_timeout=5"
```


