## Как запустить
1. Необходимо запустить сервер postgres - **sudo service postgresql start** 
2. Войти в postgres shell через postgres пользователя - **sudo -i -u postgres** 
3. Создать базу данных - **createdb tododb**
4. Запустить сервер - **go run cmd/main.go**



## Задача
To-Do List приложение, которое было создано по принципу CRUD (создание, чтение, редактирование, удаление).
Реализует REST API приложения.


## Технические требования
* Язык программирования: Golang
* БД: Postgres
* Использован фреймворк Gin (https://github.com/gin-gonic/gin) для REST запросов и запуска приложения.
* Зависимости приложения находятся в папке vendor.
* Присутствует sql-файл(migrations), где гаходятся запросы на создание таблицы.
* Все запросы проверены в приложении Postman.

## Поля сущности todo:
1) id - ID записи
2) description - Описание todo
3) created_at - Дата создания todo
4) completed_at - Дата завершения todo

## Запросы:
1) GET /todos
Показать список

2) POST /todo
Сохранение

3) PUT /todo/:id
Редактирование. Для завершения todo.

4) DELETE /todo/:id
Удаление


