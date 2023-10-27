# Сервис по созданию сокращённых ссылок

---
### Что используется?

- Golang
- gRPC
- PostgreSQL
- Docker

## Как взаимодействовать?

1. Настроить .env файл в корневой папке проекта в соответствии с локальными настройками:
    * POSTGRES_HOST
    * POSTGRES_PORT
    * POSTGRES_DB
    * POSTGRES_USER
    * POSTGRES_PASSWORD

2. Отправить запросы:

   **Создание сокращённой ссылки:**
    ```protobuf
   {   
     "url": "https://example.see"
    }
   ```

   **Возможные ответы:**
   ```json lines
    {
        "url": {
        "originalURL": "https://example.see",
        "shortenedURL": "K_OtXhzLh5"
      }
    }
   
    msg="rpc error: code = InvalidArgument ... "
   
   14 - SERVER ERROR
   ```

   **Получение оригинальной ссылки из сокращённой:**
   ```protobuf
    {
        "url": "K_OtXhzLh2"
    }
   ```

   **Возможные ответы:**
   ```json lines
    msg="getUrl success" code=0 
   
   {
      "url": {
        "originalURL": "https://example.com",
        "shortenedURL": "K_OtXhzLh5"
      }
   }
   
    msg="rpc error: code = InvalidArgument ... "
   
    msg="rpc error: code = NotFound
   
   14 - SERVER ERROR
   ```