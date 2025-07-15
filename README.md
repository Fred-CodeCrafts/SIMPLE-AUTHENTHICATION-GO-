# SIMPLE AUTHENTICATION (GO)

A simple authentication microservice built with Go, designed for modularity and scalability. Great as a base service for larger systems.

---

## ✨ Features

- 🔐 User registration with randomly generated passwords
- 🔑 Secure login using bcrypt password hashing
- 🎟️ JWT token generation on login
- 🔁 Mock password reset flow (in-memory)
- 🧱 SQLite via `modernc.org/sqlite` (pure Go, no CGO)
- 🚀 Ready for future Docker and Redis integration
- 📦 Clean project structure with separation of concerns

---

## 🧠 Tech Stack

- Go 1.21+
- Gin – web framework
- GORM – ORM
- modernc.org/sqlite – SQLite driver (pure Go)
- golang.org/x/crypto/bcrypt – password hashing
- JWT – planned (github.com/golang-jwt/jwt/v5)

---

## 📁 Project Structure

.
├── cmd/
│   └── main.go
├── controllers/
│   └── auth.go
├── models/
│   └── user.go
├── services/
│   └── auth.go
├── utils/
│   ├── hash.go
│   └── random.go
├── go.mod
├── go.sum
├── Dockerfile                # (planned)
├── docker-compose.yml        # (planned)
├── .gitattributes
├── .gitignore
├── LICENSE
└── README.md

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Git

### Clone the Repository

git clone https://github.com/Fred-CodeCrafts/SIMPLE-AUTHENTHICATION-GO-.git
cd SIMPLE-AUTHENTHICATION-GO-
go mod tidy
go run ./cmd/main.go

---

## 🔌 Sample API Endpoints

### POST /register

Request:
{
  "username": "john",
  "email": "john@example.com"
}

Response:
{
  "message": "registered",
  "password": "generated-password"
}

---

### POST /login

Request:
{
  "username": "john",
  "password": "generated-password"
}

Response:
{
  "token": "jwt-token"
}

---

## 🛣️ Roadmap (Planned Features)

This project will evolve further with:

- [ ] JWT Middleware – protect routes and validate access tokens
- [ ] Redis Integration – for token blacklist, session store, logout
- [ ] Docker + Docker Compose – containerize the service and dependencies
- [ ] Invoice Microservice – demonstrate real-world integration with auth token verification
- [ ] Swagger API Documentation – interactive docs with sample requests
- [ ] Testing – unit and integration tests for handlers and services

---

## 📄 License

This project is licensed under the MIT License.

---

## 🙋‍♂️ Author

Frederick Garner Wibowo  
Aspiring backend developer | Microservices enthusiast | Golang learner  
Indonesia
