# Веб сервер для работы с задачами
## Запуск веб сервера
- ..\go18_api\cmd\server\main.exe 
- cd ..\go18_api\cmd\server  \
go run main.go
## API
### Получение списка задач
- endpoint: GET http://localhost:8080/tasks
- код ответа: 200
- формат ответа: массив объектов в формате json
### Получение одной задачи
- GET http://localhost:8080/tasks/{id}
- код ответа: 200
- формат ответа: объект в формате json
### Создание задачи
- POST http://localhost:8080/tasks
- код ответа: 201
- формат ответа: объект в формате json
### Изменение задачи
- PUT http://localhost:8080/tasks/{id}
- код ответа: 200
- формат ответа: объект в формате json
### Удаление задачи
- DELETE http://localhost:8080/tasks/{id}
- код ответа: 204

## Тестирование работы веб сервера (с помощью CURL)
### Получение списка задач
curl -X GET http://localhost:8080/tasks -i
<br>Код ответа: 200
<br>Пример ответа: <br>
[{"id":1,"title":"Новая задача","done":false,"created_at":"2025-11-21T19:15:44+03:00"}]

### Получение одной задачи
curl -X GET http://localhost:8080/tasks/1 -i
<br>Код ответа: 200
<br>Пример ответа: <br>
{"id":1,"title":"Новая задача","done":false,"created_at":"2025-11-21T19:15:44+03:00"}

#### Ошибочный сценарий
curl -X GET http://localhost:8080/tasks/1000 -i
<br>Код ответа: 404
<br>Пример ответа: <br>
{"error":"Задача с таким ID не найдена - 1000"}

### Создание задачи
curl -X POST http://localhost:8080/tasks \
  -i \
  -H "Content-Type: application/json" \
  -d "{\"id\":0, \"title\":\"Новая задача\", \"done\":false}"
<br>Код ответа: 201
<br>Пример ответа: <br>
{"id":1,"title":"Новая задача","done":false,"created_at":"2025-11-21T19:15:44+03:00"}

#### Ошибочный сценарий
curl -X POST http://localhost:8080/tasks \
  -i \
  -H "Content-Type: application/json" \
  -d "(\"id\":0, \"title\":\"Новая задача\", \"done\":false)"
<br>Код ответа: 400
<br>Пример ответа: <br>
{"error":"Неверный формат данных"}

### Изменение задачи
curl -X PUT http://localhost:8080/tasks/1 \
  -i \
  -H "Content-Type: application/json" \
  -d "{\"id\":1, \"title\":\"Новая задача\", \"done\":true}"
<br>Код ответа: 200
<br>Пример ответа: <br>  
{"id":1,"title":"Новая задача","done":true,"created_at":"2025-11-21T19:31:34+03:00"}

#### Ошибочный сценарий
curl -X PUT http://localhost:8080/tasks/1000 \
  -i \
  -H "Content-Type: application/json" \
  -d "{\"id\":1000, \"title\":\"Новая задача\", \"done\":true}"
<br>Код ответа: 404

### Удаление задачи
curl -X DELETE http://localhost:8080/tasks/1 -i
<br>Код ответа: 204

#### Ошибочный сценарий
curl -X DELETE http://localhost:8080/tasks/1000 -i
<br>Код ответа: 500
<br>Пример ответа: <br>
{"error":"Ошибка при удалении: Элемент с таким ID не существует!"}