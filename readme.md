# Сервис синхронизации

Сервис создает клиентов, управляет статусами их алгоритмов, а также информацией о самих клиентах. 
Раз в интервал, задаваемый в конфигурационном файле, запускает либо останавливает поды согласно информации
об алгоритмах клиентов. 

### Как запустить сервис
1. Установить go, Docker и docker-сompose, если они не установлены 
2. Выполнить в консоли команду `make install-deps`
3. Выполнить команду `docker-compose up -d`
4. Выполнить команду `make migrations-up`
5. Выполнить команду `go run ./cmd/main.go`

### Как конфигурировать сервис
Конфигурация лежит в файле [config.yml](config.yml)

- db - настройка подключения к БД
- http - настройки веб-сервера
- deployWorker - настройка воркера (периодичность выполнения задается cron-строкой)

### Как делать запросы к веб-сервису

#### Создание клиента

Роутинг запроса: `POST /clients`

Пример запроса:

```
curl --location 'http://localhost:8080/clients' \
--header 'Content-Type: application/json' \
--data '{
    "client_name": "Internal",
    "version": 91,
    "image":"deposit",
    "cpu": "6th",
    "memory": "orchid",
    "priority": 276,
    "need_restart": true,
    "spawned_at" : "2024-07-12T15:20:32.641Z",
    "created_at" : "2024-07-12T15:20:32.641Z",
    "updated_at" : "2024-07-12T15:20:32.641Z"
}
'
```
#### Обновление клиента

Роутинг запроса: `POST /clients/{id}/edit`

Пример запроса:

```
curl --location 'http://localhost:8080/clients/1/edit' \
--header 'Content-Type: application/json' \
--data '{
    "client_name": "multi-byte",
    "version": 144,
    "image":"Electronics",
    "cpu": "deliverables",
    "memory": "Account",
    "priority": 105,
    "need_restart": false,
    "spawned_at" : "2024-07-12T15:27:00.582Z"
}'
```
#### Обновление статусов алгоритмов клиента

Роутинг запроса: `POST /clients/{id}/algorithmstatus`

Пример запроса:

```
curl --location 'http://localhost:8080/clients/1/algorithmstatus' \
--header 'Content-Type: application/json' \
--data '{
    "VWAP" : true,
    "TWAP" : false,
    "HFT" : false
}'
```
#### Удаление клиента

Роутинг запроса: `DELETE /clients/{id}`

Пример запроса:

```
curl --location --request DELETE 'http://localhost:8080/clients/1'
```
