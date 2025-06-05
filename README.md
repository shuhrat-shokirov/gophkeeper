# 🛡 GophKeeper

**GophKeeper** — защищённое клиент-серверное приложение для хранения чувствительных данных: логинов и паролей, текстовых
заметок, банковских карт и бинарных файлов. Сервер реализован на Go с использованием gRPC и PostgreSQL. Все данные перед
сохранением шифруются.

---

## ⚙️ Возможности

- Регистрация и вход через gRPC
- Безопасное хранение:
    - логинов и паролей
    - заметок
    - банковских карт
    - произвольных файлов
- Шифрование перед сохранением
- Взаимодействие через gRPC клиент

---

## 📦 Установка и запуск

### 🔧 Клонируй репозиторий

```bash
git clone https://github.com/shuhrat-shokirov/gophkeeper.git
cd gophkeeper
````

### 🐳 Необходимо поднять PostgreSQL сервер

```bash
docker-compose up -d
```

## 🔧 Переменные окружения

Необходимо подправить значение в `configs/config.json` либо через переменные окружения:

### Настройка grpc

Если вы хотите использовать другой адрес для gRPC сервера, то измените значение в `configs/config.json` или установите
переменную окружения `GRPC_ADDRESS`.

```json
{
  "grpc": {
    "address": "localhost:9091"
  }
}
```

### Настройка логирования

Если вы хотите изменить уровень логирования, то измените значение в `configs/config.json` или установите переменную
окружения `LOGGER_LEVEL`.

```json
{
  "logger": {
    "level": "info"
  }
}
```

### Настройка базы данных

Если у вас другой адрес базы данных, то измените значение в `configs/config.json` или установите переменную окружения
`DATABASE_DSN`.

```json
{
  "database": {
    "dsn": "postgres://postgres:postgres@localhost:5432/goph_keeper_db"
  }
}
```

### Настройка миграций

Если у вас другой адрес базы данных, то измените значение в `configs/config.json` или установите переменную окружения
`MIGRATION_DSN`.
Так же необходимо указать путь к миграциям.

```json
{
  "migration": {
    "file": "file://migrations/server",
    "dsn": "postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable"
  }
}
```

### 📧 Настройка отправки писем через Gmail

Для отправки писем необходимо указать адрес почтового сервера, логин и пароль в `configs/config.json` **или**
установить переменные окружения:

```env
EMAIL_MAIL=your_email@gmail.com
EMAIL_PASSWORD=your_app_password
```

> ⚠️ Обычный пароль от почты не подойдёт. Нужно включить двухэтапную аутентификацию и сгенерировать **пароль приложения
**.

### 🛠 Шаги по настройке:

1. **Включите двухэтапную аутентификацию:**

    * Перейдите на страницу: [https://myaccount.google.com/security](https://myaccount.google.com/security)
    * В разделе **"Вход в аккаунт Google"** нажмите **"Двухэтапная аутентификация"** и пройдите настройку.

2. **Создайте пароль приложения:**

    * Перейдите на страницу: [https://myaccount.google.com/apppasswords](https://myaccount.google.com/apppasswords)
    * Выберите **Почта** и **устройство** (или выберите «Другое» и укажите любое имя, например `GophKeeper`)
    * Скопируйте сгенерированный 16-значный пароль — он понадобится для `EMAIL_PASSWORD`

3. **SMTP-параметры для Gmail:**

   | Настройка   | Значение               |
            | ----------- | ---------------------- |
   | SMTP-сервер | `smtp.gmail.com`       |
   | Порт (SSL)  | `465`                  |
   | Порт (TLS)  | `587`                  |
   | Email       | `your_email@gmail.com` |
   | Пароль      | Пароль приложения      |

### Настройка Redis

Если вы хотите использовать Redis для кэширования, то измените значение в `configs/config.json` или установите
переменную окружения `REDIS_URL`.

```json
{
  "redis": {
    "url": "localhost:6379"
  }
}
```

### Настройка JWT

Для настройки JWT токенов, создайте пару ключей и укажите путь к приватному ключу в `configs/config.json` или установите переменную окружения `JWT_PRIVATE_KEY`.

```json
{
  "jwt": {
    "private_key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBenlLT2NkdXNHSU8vVS82Z0UxbDR0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4CklONUVVQ1A5VHhLMUUyS3NsZ0tNc0JIZ2JXSWN2RzBQaU12VFo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjUKNVZ3cm1GVWxXTEVsbk9RODJ2L3ljaVdoT2RSVEVieXE2WEFnaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGdwpzODVKZVNBWEdXNmhvRytxaHhCL2d3YmJGTlYraW11YmVGenRyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16CmJObmZjcEszb2RFZzgrM3VSWlg3QW9kdFNhOTk2T0xRTnBTSVFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG0KYXArN0R2TkU5NW5XMVFIdjM2UUczaFVUZmdGWUFSaHNDazkwcndJREFRQUJBb0lCQUZ1Wkx2L1h3b2VHdjNYUAovSThCK25rYTJEUkM1Mm5SMnluSzVYa01lWlI1bDQ0dDl3Zzc4bDYwUTZFVHAwSysySTV2NnpJemZaa3hrWDZjCkJnb2JCTTVlQUIxQ1pqTUFnQnZqRUpxd21qV3ArWitNSkhtU3BHNjFmSkgzcUtmMFlKc0NJMzlwOTUzQXNNbC8Kc3Q4UGFEdklQcjNOT29ITTdxSjc4UnYvejkvRVNSdVBuUFZPaEtQSXh6OFZ4Z2p5eGRlemM5Y0NvbnZLbzV0bgpXUGFjOStlYkwvSTFFWDdHZ1IwVnZzbDdtc0RoVVorcnkraUdnM3pBVzhIOGJ5NUFGR3lkdUNjeGtrWmVUVFRFClR1SllVSGdCTWN5NUtCNStxdzNvRk1QVnFNeTE3RktCRncwMW8vQlE3WlFCdERNUFhoWVJyY2RDRk9BOThia1QKbzlnQzFka0NnWUVBNTVxZ0NUUzFibk9xY3lmOFpMdkM2ZFdyVWQ1Q1VZVEZiNWZWUGlYRmxpTTExR2FRbWxGUgpyTmxjTU83ZEVHNVNvUmkxNzJHQVJZK24xZi84QXlGN2g0WHlqanFkbnNybmhFckFJRXNiN1h1cFRzM3RqL1pXCjhJYWFFZ0VxUExITHVrdHhyWjh2aHdtamxlSXhrVmVGR0N1S3JSMFc0bWRlV3ArNmJjYnJ2TmtDZ1lFQTVQUWIKSHI4c2FhdTIrSHVTTm90czl5ZUhNdmRKczkvZzdmeHVzUGs4STRrVHdGZHpZeUhNSmcvdzEwTTY3VFNiYzJFcAp1SXlxSlpCVUNDNkVleUF3cjJERjBiRkJEWGpqZ3hFQzYrMU1QdXlXV2JWam1EaWRreFA4WVRkb0k5ZjVrbklXClBxcUhSMWdEY3lJVW14UE5UOXVqZTBBQXJWZXRZZUp4bmU1ajZNY0NnWUVBby84RGIwRlpiMXFMeVhyNDUwTmsKNHpzZlVwczFEcEFiVmNlSGdiZ3hUdnlqczBEbEI4Q3BPdUcydkJlSGhZajVEWVYzM29lRjBydkVObTVLdnRUSQpxZEFaVHNrR3IxZ3gwNlV5b2l0TkhUNUJSc0hlZytBRTg4Lzc3Ti9TVHFQL0JHMURrNU55amdZdlJZU2pZSzArCld6MEp0MGN2MnlVaTFMemh2N1hwV3hFQ2dZQWNVSmdlQ0ZTTXlRQzY0RVZuMjN4aFlKRVcyNEJRNzRvWXhKUkgKN0xya1JpcWNLZlNLT1A3UFlqOU56L0cwcmtIZlZnL2IxQUdpM2FPVzAzSHM3RUU1SDBXM3RpMHVabG4wdHFEZQozcDBFVnl3TThpTGNDM3hwV1Jwb1Izcm9tK2d3bFUxcytKZjhXY1VyY3ZhTGF6cUQrc3pRREUxSklzTzlqRXl5CjFHMmt0d0tCZ0FhdjY0UlhjWFM0SGRrK1JJMUlmSTdkR0h1b3FQRjB4ZVhzcnZCeTNWeUxpK1RHRzNraXNMMFkKY1F4S1RxODdUckpSWEE5YjFYU0ZHV0VvenVOQ2VCZjNEMGhTTWpkcXg2SEgrQ0MrTmdNOGtZUDBycWQ0N1plbAp3UzRyVWdpbm91TlFZYVlRaEpLQVd0Y21JN0JLNGhBb1JOLzRubE9hSXZDa2dPWGVpRW8zCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
  }
}
```

### Настройка AES
Если вы хотите использовать другой ключ для шифрования, то измените значение в `configs/config.json` или установите
переменную окружения `AES_SECRET_KEY`.

```json
{
  "aes": {
    "secret_key": "1234567890abcdef"
  }
}
```

### Запуск сервера
Для запуска сервера выполните команду:

```bash
go run cmd/server/main.go
```

### Запуск клиента
Для запуска клиента необходимо установить переменную окружения `GOPH_KEEPER_GRPC_ADDRESS` с адресом gRPC сервера.
И так же необходимо задать публичный ключ JWT в переменной окружения `GOPH_KEEPER_PUBLIC_KEY`.

Для запуска клиента выполните команду:

```bash
export GOPH_KEEPER_GRPC_ADDRESS=localhost:9091
export GOPH_KEEPER_PUBLIC_KEY="LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF6eUtPY2R1c0dJTy9VLzZnRTFsNAp0Z3Z4KzR5TUExQ1pQRkYwRnhKbzJaS1ZTY3I4SU41RVVDUDlUeEsxRTJLc2xnS01zQkhnYldJY3ZHMFBpTXZUClo0dTB3SWlPaTVSMDNlK3I5V1NqOG1xSCs3UjU1VndybUZVbFdMRWxuT1E4MnYveWNpV2hPZFJURWJ5cTZYQWcKaU5BckZyL3NFRTBacHFPdlVSeEFmeG5Qb1ZGd3M4NUplU0FYR1c2aG9HK3FoeEIvZ3diYkZOVitpbXViZUZ6dApyd1NJUGdmNjR2d2RoWnpDY1JZOVRUK1dyRm16Yk5uZmNwSzNvZEVnOCszdVJaWDdBb2R0U2E5OTZPTFFOcFNJClFnUDRiMnBXem5hc0NlRHU3ZlBWME5GNkJmSG1hcCs3RHZORTk1blcxUUh2MzZRRzNoVVRmZ0ZZQVJoc0NrOTAKcndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="

go run cmd/client/main.go
```
