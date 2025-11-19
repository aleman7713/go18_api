# Веб сервер для работы с задачами
## Запуск веб сервера
- ..\go18_api\cmd\server\main.exe 
- cd ..\go18_api\cmd\server  \
go run main.go
## API
### Получение списка задач
- GET http://localhost:8080/tasks
- POST http://localhost:8080/tasks
### Получение одной задачи
- GET http://localhost:8080/tasks/{id}
### Создание задачи
- PUT http://localhost:8080/tasks/
### Изменение задачи
- PUT http://localhost:8080/tasks/{id}
### Удаление задачи
- DELETE http://localhost:8080/tasks/{id}
## Тестирование работы веб сервера (с помощью CURL)
### Получение списка задач
curl -X GET http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks
### Получение одной задачи
curl -X GET http://localhost:8080/tasks/1
### Создание задачи
curl -X PUT http://localhost:8080/tasks/ \
  -H "Content-Type: application/json" \
  -d "{\"id\":0, \"title\":\"Новая задача\", \"done\":false}"
### Изменение задачи
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d "{\"id\":1, \"title\":\"Новая задача\", \"done\":true}"
### Удаление задачи
curl -X DELETE http://localhost:8080/tasks/1