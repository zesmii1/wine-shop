Wine Shop

Проект представляет собой REST API для управления магазином вин. Реализован на Go с использованием Gin, GORM и PostgreSQL. Поддерживает регистрацию, аутентификацию, роли пользователей, CRUD для вин, а также контейнеризацию через Docker и docker-compose.

Старт проекта

1. Клонировать репозиторий

git clone https://github.com/yourusername/wine-shop.git
cd wine-shop

2. Запуск через Docker Compose

docker-compose up --build

3. Сервер будет доступен по адресу:

http://localhost:8080

Аутентификация

Используется JWT-токен с ролями. Роль вшивается в токен при логине.

Пример токена: Bearer eyJhbGci...

Postman Collection

Для удобного тестирования API можно использовать коллекцию Postman.

Файл: Wine Shop.postman_collection.json

Импортируйте его в Postman и тестируйте следующие запросы:

POST /api/auth/register

POST /api/auth/login

GET /api/users

DELETE /api/users/:id

GET /api/v1/wines

POST /api/v1/wines

GET /api/v1/wines/:id

PUT /api/v1/wines/:id

DELETE /api/v1/wines/:id

GET /api/v1/me

Структура проекта

wine-shop/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
├── internal/
│   ├── auth/        # Авторизация и JWT
│   ├── db/          # Подключение к БД
│   ├── models/      # Модели данных
│   └── routes/      # Маршруты и middleware
└── Wine Shop.postman_collection.json

🛠 Используемые технологии

Go (Gin, GORM)

PostgreSQL

Docker + Docker Compose

JWT аутентификация

Автор

Қалтай Сезім, 2025