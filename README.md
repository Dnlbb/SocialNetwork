# Социальная сеть с функциями регистрации, аутентификации и авторизации на основе JWT

## Описание

Этот проект представляет собой базовую социальную сеть, в которой пользователи могут регистрироваться, входить в систему и просматривать свои профили. В качестве механизма авторизации используется JSON Web Tokens (JWT).

## Функционал

- Регистрация пользователей
- Вход пользователей
- Защищенные маршруты с использованием JWT
- Получение информации о профиле пользователя

## Структура проекта

```plaintext
SocialNetwork/
├── go.mod
├── go.sum
├── main.go
├── README.md
├── handlers/
│   ├── auth.go
│   └── user.go
├── utils/
│   ├── auth.go
│   └── jwt.go
└── models/
    └── user.go
```

## Использование

- Регистрация: `POST /register`
- Логин: `POST /login`
- Профиль: `GET /profile` (требуется авторизация)

## Зависимости

- [Fiber](https://github.com/gofiber/fiber) - веб-фреймворк
- [Logrus](https://github.com/sirupsen/logrus) - логгер
- [JWT Middleware](https://github.com/gofiber/contrib/tree/main/jwt) - middleware для работы с JWT

