#  Test Task

## Installation
1)git clone https://github.com/evsyukoov/crud_task.git

2)sh entrypoint.sh

3)Развернется 2 контейнера, в одном сервер, во втором MySQL. Сервер слушает на 80 порту.
Для доступа к MySQL c хост машины и проверки добавленных данных нужно стучаться на 127.0.0.1:12345, admin 1111

## API

### Метод показа статистики

Метод: GET

1) /show?from=YYYY-MM-DD&to=YYYY-MM-DD - отдаст все данные отсортированные по дате
   
2) /show/clicks?from=YYYY-MM-DD&to=YYYY-MM-DD - отдаст все данные отсортированные по количеству кликов
   
3) /show/views?from=YYYY-MM-DD&to=YYYY-MM-DD - отдаст все данные отсортированные по количеству просмотров
   
4) /show/cost?from=YYYY-MM-DD&to=YYYY-MM-DD - отдаст все данные отсортированные по цене

#### Примеры использования

curl -v  "localhost/show?from=2000-01-02&to=2020-02-02"

curl -v  "localhost/show/views?from=2000-01-02&to=2020-02-02"

#### Коды ответа

1) 200 OK - запрос обработан успешно
2) 400 Bad Request - Ошибка пути или query-string
3) 500  InternalServerError - Ошибка соединения с БД

### Метод добавления статистики

Метод: POST

Путь: /save
Тело:  Тело в формате JSON '{"date":"YYYY-MM_DD", "clicks":4, "views":6, "cost":10.21}'

#### Примеры использования

curl -v -X POST --data '{"date":"2020-01-02","views":2,"clicks":3,"cost":1.11}'

curl -v -X POST --data '{"date":"2020-01-02"}'

curl -v -X POST --data '{"date":"2020-01-02", "views":2}'

#### Коды ответа

1) 200 OK - добавление прошло успешно
2) 400 Bad Request - Ошибка пути или формата тела
3) 500  InternalServerError - Ошибка соединения с БД
   

### Метод удаления

Метод: DELETE

Route: /clear

#### Коды ответа

1) 200 OK - запрос обработан успешно
2) 400 Bad Request - Ошибка пути
3) 500  InternalServerError - Ошибка соединения с БД