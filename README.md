# 📚 Library API – Golang RESTful Service

A simple **Golang REST API** for managing Books, Categories, and Users — built with the **Gin** framework, **PostgreSQL**, and **JWT authentication**.

---

## 🚀 Features

- ✅ User authentication with JWT  
- ✅ CRUD for Categories and Books  
- ✅ Auto-detects and runs SQL migrations from `db/sql_migrations/migrate.sql`  
- ✅ Input validation for requests  
- ✅ Environment-based configuration via `.env`  
- ✅ Lightweight and Railway-compatible  

---

## 🏗 Tech Stack

| Component | Technology |
|------------|-------------|
| Language | Go 1.21+ |
| Web Framework | [Gin Gonic](https://github.com/gin-gonic/gin) |
| Database | PostgreSQL |
| Auth | JWT (HS256) |
| ORM | Native SQL (no GORM) |
| Migrations | SQL scripts |
| Environment | [godotenv](https://github.com/joho/godotenv) |

---

## 📁 Folder Structure

```bash
library/
├── main.go
│
├── config/
│   ├── .env
│   └── config.go
│
├── controllers/
│   ├── auth_controller.go
│   ├── books_controller.go
│   ├── categories_controller.go
│   └── jwt.go
│
├── db/
│   ├── db.go
│   └── sql_migrations/
│       └── migrate.sql
│
├── repository/
│   ├── users_repository.go
│   ├── categories_repository.go
│   └── books_repository.go
│
└── structs/
    └── models.go
```

---

## ⚙️ Setup Instructions

### 1️⃣ Clone and install dependencies
```bash
git clone https://github.com/yourusername/library-api.git
cd library-api
go mod tidy
```

---

### 2️⃣ Create `.env` file under `config/`

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=ggwp
DB_NAME=postgres

PORT=8080
JWT_SECRET=myjwtsecret
```

💡 For production, generate a strong secret key:
```bash
openssl rand -base64 64
```

---

### 3️⃣ Run the application
```bash
go run main.go
```

Expected output:
```
✅ Loaded configuration: DB_HOST=localhost, DB_NAME=postgres
🔌 Connecting to database...
🚀 Running migration from migrate.sql ...
✅ Migration executed successfully
🚀 Server running on port 8080
```

---

## 🔐 Authentication

A default admin user is automatically created during migration:

| Username | Password |
|-----------|-----------|
| admin | admin123 |

### 🔑 Login to get JWT token

```bash
curl -X POST http://localhost:8080/api/users/login   -H "Content-Type: application/json"   -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

**Response**
```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5..." }
```

Save your token:
```bash
export TOKEN=<paste_your_token_here>
```

---

## 🏷️ Categories API

### 🟢 Get All Categories
```bash
curl -X GET http://localhost:8080/api/categories   -H "Authorization: Bearer $TOKEN"
```

---

### 🟡 Create Category
```bash
curl -X POST http://localhost:8080/api/categories   -H "Content-Type: application/json"   -H "Authorization: Bearer $TOKEN"   -d '{"name": "Fiction"}'
```

**Response**
```json
{ "message": "category created" }
```

---

### 🔵 Get Category by ID
```bash
curl -X GET http://localhost:8080/api/categories/1   -H "Authorization: Bearer $TOKEN"
```

---

### 🔴 Delete Category
```bash
curl -X DELETE http://localhost:8080/api/categories/1   -H "Authorization: Bearer $TOKEN"
```

**Response**
```json
{ "message": "category deleted" }
```

---

### 📘 Get Books by Category
```bash
curl -X GET http://localhost:8080/api/categories/2/books   -H "Authorization: Bearer $TOKEN"
```

**Response**
```json
{
  "data": [
    {
      "id": 2,
      "title": "Clean Code",
      "release_year": 2020,
      "price": 250000,
      "total_page": 320,
      "thickness": "tebal"
    }
  ]
}
```

---

## 📚 Books API

### 🟢 Get All Books
```bash
curl -X GET http://localhost:8080/api/books   -H "Authorization: Bearer $TOKEN"
```

---

### 🟡 Create New Book
```bash
curl -X POST http://localhost:8080/api/books   -H "Content-Type: application/json"   -H "Authorization: Bearer $TOKEN"   -d '{
    "title": "Clean Code",
    "description": "A handbook of agile software craftsmanship",
    "image_url": "https://example.com/cleancode.jpg",
    "release_year": 2020,
    "price": 250000,
    "total_page": 320,
    "category_id": 2
  }'
```

**Response**
```json
{
  "message": "book created",
  "thickness": "tebal"
}
```

---

### 🔵 Get Book by ID
```bash
curl -X GET http://localhost:8080/api/books/2   -H "Authorization: Bearer $TOKEN"
```

**Response**
```json
{
  "data": {
    "id": 2,
    "title": "Clean Code",
    "release_year": 2020,
    "price": 250000,
    "total_page": 320,
    "thickness": "tebal"
  }
}
```

---

### 🔴 Delete Book
```bash
curl -X DELETE http://localhost:8080/api/books/2   -H "Authorization: Bearer $TOKEN"
```

**Response**
```json
{ "message": "book deleted" }
```

---

## 🧠 API Summary

| Method | Endpoint | Description |
|--------|-----------|-------------|
| **POST** | `/api/users/login` | Get JWT token (login) |
| **GET** | `/api/categories` | List all categories |
| **POST** | `/api/categories` | Create new category |
| **GET** | `/api/categories/:id` | Get category detail |
| **DELETE** | `/api/categories/:id` | Delete category |
| **GET** | `/api/categories/:id/books` | List books in a category |
| **GET** | `/api/books` | List all books |
| **POST** | `/api/books` | Create new book |
| **GET** | `/api/books/:id` | Get book details |
| **DELETE** | `/api/books/:id` | Delete a book |

---

## 💾 Default Data from Migration

| Type | Name / Value |
|------|---------------|
| Default user | `admin` / `admin123` |
| Default categories | `Technology`, `Science` |
| Default DB name | `postgres` |

---

## 👨‍💻 Author

**Gonewaje**  
DevOps Engineer • Backend Developer  
🌐 [https://www.gonewaje.cloud](https://www.gonewaje.cloud)

---