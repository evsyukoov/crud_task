#  Test Task

## Установка в докере

1)git clone https://github.com/evsyukoov/crud_task.git

2)cd crud_task/Avito && sh entrypoint.sh

3)Развернется 2 контейнера, в одном сервер, во втором MySQL. Сервер слушает на 80 порту.
Для доступа к MySQL c хост машины и проверки добавленных данных нужно стучаться на 127.0.0.1:12345, admin 1111
Команда - mysql -h 127.0.0.1 -u admin -P12345 -p1111

##  Установка на своей машине

1)Создаем БД с именем avito_test.  Все необходимые скрипты в файле scrypt.sql

2)В файле ./crud_task/Avito/Application/config/cofig.json устанавливаем все необходимые параметры для соединения с БД

3)cd ./crud_task/Avito/Application && make run

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

####Примечание

В пришедшем Json в ключах cpc/cpm значения равные -1 обозначают clicks/views = 0 в БД (поскольку в методе добавления

сказано про опциональность этих полей 

#### Коды ответа

1) 200 OK - запрос обработан успешно
2) 400 Bad Request - Ошибка пути или query-string
3) 500  InternalServerError - Ошибка соединения с БД

### Метод добавления статистики

Метод: POST

Путь: /save

Тело:  Тело в формате JSON '{"date":"YYYY-MM_DD", "clicks":4, "views":6, "cost":10.21}'

#### Примеры использования

curl -v -X POST localhost/save --data '{"date":"2020-01-02","views":2,"clicks":3,"cost":1.11}'

curl -v -X POST localhost/save --data '{"date":"2020-01-02"}'

curl -v -X POST localhost/save --data '{"date":"2020-01-02", "views":2}'

#### Коды ответа

1) 200 OK - добавление прошло успешно
2) 400 Bad Request - Ошибка пути или формата тела
3) 500  InternalServerError - Ошибка соединения с БД
   

### Метод удаления

Метод: DELETE

Route: /clear

#### Примеры использования

curl -v -X DELETE localhost/clear

#### Коды ответа

1) 204 OK - запрос обработан успешно
2) 400 Bad Request - Ошибка пути
3) 500  InternalServerError - Ошибка соединения с БД