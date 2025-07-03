# Техническое задание: Переводчик с использованием OpenRouter API

## Описание проекта

Веб-сервис для перевода текста с использованием нейросетей через OpenRouter API. Сервис принимает текст на одном языке и возвращает перевод на другой язык.

## Архитектура приложения

### Структура проекта

```
/translator-service
├── main.go
├── internal/
│   ├── handlers/
│   │   └── translator_handler.go
│   ├── services/
│   │   └── translator_service.go
│   ├── client/
│   │   └── openrouter_client.go
│   ├── entity/
│   │   └── translation.go
│   └── dto/
│       ├── request.go
│       └── response.go
├── config/
│   └── config.go
└── go.mod
```

### Слои и их ответственности

#### 1. Handlers (`internal/handlers/`)
**Файл:** `translator_handler.go`

**Ответственность:**
- Обработка HTTP запросов через Gin
- Валидация входящих данных
- Вызов сервисного слоя
- Формирование HTTP ответов

**Ограничения:**
- НЕ содержит бизнес-логику
- НЕ знает детали работы с внешними API
- Работает только с DTO структурами

**Основные методы:**
- `TranslateText(c *gin.Context)` - основной хендлер для перевода
- `GetSupportedLanguages(c *gin.Context)` - получение списка поддерживаемых языков

#### 2. Services (`internal/services/`)
**Файл:** `translator_service.go`

**Ответственность:**
- Бизнес-логика приложения
- Обработка и валидация данных
- Вызов клиента для работы с OpenRouter
- Преобразование данных между слоями

**Ограничения:**
- НЕ работает с gin.Context
- НЕ знает про HTTP статус-коды
- Возвращает ошибки в виде Go errors

**Основные методы:**
- `TranslateText(text, fromLang, toLang string) (*entity.Translation, error)`
- `GetSupportedLanguages() ([]string, error)`

#### 3. Client (`internal/client/`)
**Файл:** `openrouter_client.go`

**Ответственность:**
- HTTP запросы к OpenRouter API
- Сериализация/десериализация JSON
- Обработка HTTP ошибок и таймаутов
- Управление API ключами

**Ограничения:**
- НЕ содержит бизнес-логику
- Работает только с низкоуровневыми HTTP запросами
- Возвращает "сырые" данные от API

**Основные методы:**
- `SendTranslationRequest(prompt string) (*OpenRouterResponse, error)`
- `HealthCheck() error`

#### 4. Entity (`internal/entity/`)
**Файл:** `translation.go`

**Ответственность:**
- Доменные сущности
- Бизнес-объекты приложения
- Чистые структуры без зависимостей

**Структуры:**
```go
type Translation struct {
    ID           string
    OriginalText string
    TranslatedText string
    FromLanguage string
    ToLanguage   string
    CreatedAt    time.Time
}
```

#### 5. DTO (`internal/dto/`)
**Файлы:** `request.go`, `response.go`

**Ответственность:**
- Структуры для передачи данных между слоями
- JSON теги для сериализации
- Валидация входящих данных

**Структуры:**
```go
// request.go
type TranslateRequest struct {
    Text         string `json:"text" binding:"required"`
    FromLanguage string `json:"from_language" binding:"required"`
    ToLanguage   string `json:"to_language" binding:"required"`
}

// response.go
type TranslateResponse struct {
    TranslatedText string `json:"translated_text"`
    FromLanguage   string `json:"from_language"`
    ToLanguage     string `json:"to_language"`
    Success        bool   `json:"success"`
}
```

#### 6. Config (`config/`)
**Файл:** `config.go`

**Ответственность:**
- Конфигурация приложения
- Загрузка переменных окружения
- Настройки API и сервера

**Параметры:**
- OpenRouter API ключ
- URL эндпоинтов
- Таймауты
- Порт сервера

## API Endpoints

### POST /translate
Перевод текста

**Запрос:**
```json
{
    "text": "Hello world",
    "from_language": "en",
    "to_language": "ru"
}
```

**Ответ:**
```json
{
    "translated_text": "Привет мир",
    "from_language": "en",
    "to_language": "ru",
    "success": true
}
```

### GET /languages
Получение списка поддерживаемых языков

**Ответ:**
```json
{
    "languages": ["en", "ru", "de", "fr", "es"],
    "success": true
}
```

### GET /health
Проверка состояния сервиса

**Ответ:**
```json
{
    "status": "ok",
    "timestamp": "2025-07-04T10:30:00Z"
}
```

## Поток данных

1. **Клиент** → POST /translate → **Handler**
2. **Handler** → валидация → **Service**
3. **Service** → бизнес-логика → **Client**
4. **Client** → HTTP запрос → **OpenRouter API**
5. **OpenRouter API** → ответ → **Client**
6. **Client** → сырые данные → **Service**
7. **Service** → обработка → **Handler**
8. **Handler** → JSON ответ → **Клиент**

## Технические требования

### Зависимости
- `github.com/gin-gonic/gin` - веб-фреймворк
- `net/http` - HTTP клиент (стандартная библиотека)
- `encoding/json` - JSON сериализация
- `os` - переменные окружения

### Переменные окружения
- `OPENROUTER_API_KEY` - API ключ для OpenRouter
- `OPENROUTER_BASE_URL` - базовый URL API (по умолчанию: https://openrouter.ai/api/v1)
- `SERVER_PORT` - порт сервера (по умолчанию: 8080)
- `HTTP_TIMEOUT` - таймаут HTTP запросов (по умолчанию: 30s)

### Обработка ошибок
- Валидация входных данных
- Обработка сетевых ошибок
- Логирование ошибок
- Graceful shutdown

## Этапы разработки

1. **Этап 1:** Создание структуры проекта и базовых файлов
2. **Этап 2:** Реализация DTO и Entity
3. **Этап 3:** Разработка OpenRouter клиента
4. **Этап 4:** Создание сервисного слоя
5. **Этап 5:** Реализация handlers с Gin
6. **Этап 6:** Настройка конфигурации
7. **Этап 7:** Тестирование и отладка

## Примечания

- Приложение stateless - не требует БД
- Используется чистая архитектура с разделением ответственности
- Код должен быть легко тестируемым
- Возможность легкого добавления новых AI провайдеров