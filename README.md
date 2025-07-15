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

## 🚀 Getting Started

### Clone the repo

```bash
git clone https://github.com/Fred-CodeCrafts/SIMPLE-AUTHENTHICATION-GO-.git
cd SIMPLE-AUTHENTHICATION-GO-
go mod tidy
go run ./cmd/main.go

