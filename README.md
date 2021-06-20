# url-shortener

## Как запустить 
Создать таблицу в Postgres:

```sql
CREATE TABLE links (
    id SERIAL PRIMARY KEY ,
    url TEXT NOT NULL UNIQUE 
)
```

Настройка происходит с помощью dev.env файла в /config (пример - example.dev.env).

## API

### Примеры

Получить сокращенную ссылку:

```http request
POST localhost:3001/short
Content-Type: application/json

{
  "url": "https://vk.com/"
}
```

Ответ:
```json
{
  "url": "https://igor.r/BM"
}
```

##

Получить полную ссылку

```http request
POST localhost:3001/long
Content-Type: application/json

{
  "url": "https://igor.r/Bk"
}
```

Ответ:
```json
{
  "url": "https://www.ozon.ru/"
}
```
