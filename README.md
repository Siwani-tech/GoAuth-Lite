# 🔐 GoAuth Lite

GoAuth Lite is a lightweight authentication service built in Go (Golang).  
It provides a simple and clean implementation of user authentication using JWT (JSON Web Tokens).

---

## ✨ Features

- ✅ User Signup with password hashing (bcrypt)
- 🔐 Secure Login with JWT token generation
- 🛡️ Protected routes using middleware
- ⚡ Lightweight & fast (uses Go's standard net/http)
- 🧩 Clean architecture (handlers, services, repository, middleware)

---



## 🚀 Getting Started

### 1️⃣ Clone the repository

```bash
git clone https://github.com/Siwani-tech/GoAuth-Lite.git
cd GoAuth-Lite
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Run the server

```bash
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### 🟢 Health Check

```
GET /health
```

Response:
```
GoAuth Lite is running
```

---

### 📝 Signup

```
POST /signup
```

Body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

---

### 🔑 Login

```
POST /login
```

Body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "token": "your-jwt-token"
}
```

---

### 👤 Profile (Protected)

```
GET /profile
```

Headers:
```
Authorization: Bearer <your-token>
```

Response:
```
Hi there user@example.com
```

---

## 🔐 Authentication Flow

1. User signs up → password is hashed using bcrypt  
2. User logs in → receives JWT token  
3. Token is sent in `Authorization` header  
4. Middleware validates token and extracts user email  
5. Protected routes use the authenticated context  

---

## 🛠️ Tech Stack

- Go (Golang)
- net/http
- bcrypt (password hashing)
- JWT (github.com/golang-jwt/jwt/v5)

---


Built with ❤️ using Go
