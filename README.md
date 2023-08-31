# Интрукция по запуску
1. Запустить контейнер с базой данных<br>
`docker compose run db`
2. Запустить контейнер с сервером<br>
`docker compose run avito-app`

# Эксплуатация
OpenAPI документация лежит в /api

## Основное задание
1. Создание пользователя<br>
`curl --location 'localhost:8080/user/add' \
   --header 'Content-Type: text/plain' \
   --data '{
   "id": 5
   }'`
2. Удаление пользователя<br>
`curl --location 'localhost:8080/user/remove' \
   --header 'Content-Type: text/plain' \
   --data '{
   "id": 5
   }'`
3. Создание сегмента<br>
`curl --location 'localhost:8080/segment/add' \
   --header 'Content-Type: text/plain' \
   --data '{
   "slug": "segment_4"
   }'`
4. Удаление сегмента<br>
   `curl --location 'localhost:8080/segment/remove' \
   --header 'Content-Type: text/plain' \
   --data '{
   "slug": "segment_4"
   }'`
5. Обновление сегментов пользователя<br>
`curl --location 'localhost:8080/usersegments/update' \
   --header 'Content-Type: text/plain' \
   --data '{
   "user_id": 5,
   "add_slugs": [
   "segment_5",
   "segment_6"
   ],
   "delete_slugs": [
   "segment_4",
   "segment_2"
   ]
   }'`
6. Получение активных сегментов пользователя<br>
`curl --location 'localhost:8080/usersegments/get' \
   --header 'Content-Type: text/plain' \
   --data '{
   "id": 5
   }'`

## Дополнительное задание 3
1. Создание сегмента с изначальными пользователями<br>
`curl --location 'localhost:8080/segment/add' \
   --header 'Content-Type: text/plain' \
   --data '{
   "slug": "segment_4",
   "percent": 50
   }'`

