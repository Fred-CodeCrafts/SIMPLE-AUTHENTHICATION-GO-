# SIMPLE AUTHENTICATION (GO)

A simple authentication microservice in Go designed to be modular and extensible.

## ✨ Features

- 🔐 Register + Login with bcrypt password hashing
- 🔁 Auto-generated password upon registration
- 🧪 Mock password reset flow
- 🔑 JWT token generation
- 🧱 SQLite via modernc (pure Go – no CGO!)
- ⚙️ Ready for Redis and Docker integration
- 🧩 Modular structure (controller/service/model/util)

## 🧠 Tech Stack

- Go 1.21+
- Gin Web Framework
- GORM ORM
- modernc.org/sqlite
- JWT
- bcrypt

## 📁 Project Structure

.
├── cmd/
│ └── main.go
├── controllers/
│ └── auth.go
├── models/
│ └── user.go
├── services/
│ └── auth.go
├── utils/
│ ├── hash.go
│ └── random.go
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── README.md
└── LICENSE

## 🚀 Getting Started

### Clone the repo

```bash
git clone https://github.com/Fred-CodeCrafts/SIMPLE-AUTHENTHICATION-GO-.git
cd SIMPLE-AUTHENTHICATION-GO-
go mod tidy
go run ./cmd/main.go

Sample API
POST /register
json
Copy
Edit
{
  "username": "john",
  "email": "john@example.com"
}
Returns:

json
Copy
Edit
{
  "message": "registered",
  "password": "random-password"
}
POST /login
json
Copy
Edit
{
  "username": "john",
  "password": "random-password"
}
🧪 To Do (Planned features)
 JWT verification middleware

 Redis token blacklist

 Email support for password reset

 Docker Compose integration

 Swagger API docs
