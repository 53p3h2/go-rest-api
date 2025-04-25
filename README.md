# 📝 Task API with Gin + SQLite

A minimal REST API built using [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/), and [SQLite3](https://www.sqlite.org/index.html). This app handles basic CRUD operations for a `Task` model.

---

## 📌 Notes
- This project is meant for university course learning project.
- It runs fully offline using SQLite.
- due to sqlite3 driver needing gcc it only runs on unix systems with gcc installed

---

## 📁 Project Setup

### ✅ Clone the repository
```bash
git clone https://github.com/53p3h2/go-rest-api.git
cd go-rest-api
```

### ✅ Install dependencies
Make sure you have Go installed (version 1.18+). Then run:
```bash
go mod tidy
```

### ✅ Run the application
```bash
go run main.go
```
The API will be available at: `http://localhost:3001`

---

## 🧱 Database

- A `tasks.db` file will be created automatically in the project root.
- It uses SQLite and is managed by GORM.

---

## 📦 Task Model
```json
{
  "id": 1,
  "text": "example task",
  "status": true
}
```
- `id`: auto-incremented integer
- `text`: string
- `status`: boolean (true = completed, false = incomplete)

---

## 🚀 API Endpoints and Usage (with `curl`)

### 🔹 GET `/tasks` — Get all tasks
```bash
curl -X GET http://localhost:3001/tasks
```
**Response:**
```json
[
  { "id": 1, "text": "Learn Go", "status": false },
  { "id": 2, "text": "Build API", "status": true }
]
```

---

### 🔹 POST `/tasks` — Create a new task
```bash
curl -X POST http://localhost:3001/tasks \
  -H "Content-Type: application/json" \
  -d '{"text": "Read book", "status": false}'
```
**Response:**
```json
{ "id": 3, "text": "Read book", "status": false }
```

---

### 🔹 PUT `/tasks/:id` — Update a task
```bash
curl -X PUT http://localhost:3001/tasks/3 \
  -H "Content-Type: application/json" \
  -d '{"text": "Read book", "status": true}'
```
**Response:**
```json
{ "id": 3, "text": "Read book", "status": true }
```

---

### 🔹 DELETE `/tasks/:id` — Delete a task
```bash
curl -X DELETE http://localhost:3001/tasks/1
```
**Response:**
```json
{ "message": "Task deleted" }
```

---


## 🛠 Tools Used
- [Gin](https://github.com/gin-gonic/gin) — HTTP Web Framework
- [GORM](https://gorm.io/) — ORM for Go
- [SQLite3](https://www.sqlite.org/index.html) — Lightweight embedded database

---

## 📬 License
MIT — free to use and modify.

