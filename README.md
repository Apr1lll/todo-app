# ✅ Todo App на Go

Простое Todo-приложение с REST API на Go и веб-интерфейсом.

## 🚀 Возможности

- ✨ Создание задач с заголовком и описанием
- 📋 Просмотр всех задач
- ✅ Отметка задач как выполненных
- 🔄 Возобновление задач
- 🗑️ Удаление задач
- 🔍 Фильтрация (все/активные/завершенные)
- 📅 Автоматическая фиксация времени создания и завершения

## 🛠️ Технологии

- **Backend**: Go 1.24+
- **Router**: Gorilla Mux
- **Frontend**: HTML5, CSS3, Vanilla JavaScript

## 📦 Установка и запуск

### Предварительные требования

- Go 1.24 или выше
- Git (для клонирования репозитория)

### Локальный запуск

1. **Клонируйте репозиторий**

```bash
git clone https://github.com/Apr1lll/todo-app.git
cd todo-app
```

2. **Запусти бэкенд**

```bash
go mod download
go run main.go
```

Бэкенд запустится на http://localhost:9091

3. **Запусти фронтенд**
   Самый простой способ - используй Live Server в VS Code:

   -Установи расширение "Live Server"

   -Открой файл frontend/index.html

   -Нажми "Go Live" в правом нижнем углу

Или через Python:

```bash
cd frontend
python3 -m http.server 8080
```

4. **Открой приложение**

Перейди по адресу: http://localhost:8080

**📋 API Endpoints**
Метод URL Описание
GET /tasks Получить все задачи
GET /tasks?completed=false Получить незавершенные задачи
GET /tasks/{title} Получить задачу по названию
POST /tasks Создать новую задачу
PATCH /tasks/{title} Изменить статус задачи
DELETE /tasks/{title} Удалить задачу

**🔍 Примеры запросов**
-Создать задачу

```bash
curl -X POST http://localhost:9091/tasks \
-H "Content-Type: application/json" \
-d '{"Title":"Купить хлеб","Description":"В магазине у дома"}'
```

-Получить все задачи
bash

```bash
curl http://localhost:9091/tasks
```

-Отметить задачу как выполненную

```bash
curl -X PATCH http://localhost:9091/tasks/Купить%20хлеб \
-H "Content-Type: application/json" \
-d '{"Complete":true}'
```

-Удалить задачу

```bash
curl -X DELETE http://localhost:9091/tasks/Купить%20хлеб
```

5. **📄 Лицензия**

MIT License. Используй для обучения и своих проектов.

6. **👨‍💻 Автор**

GitHub: @Apr1lll
