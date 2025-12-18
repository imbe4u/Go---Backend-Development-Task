User Management REST API (Go + Fiber + SQLC)

A clean, production-ready RESTful API built using GoFiber, PostgreSQL, and SQLC to manage users with name and date of birth, while calculating age dynamically using Go’s time package.

This project follows a layered architecture and industry-standard directory structure, making it suitable for real-world applications and interviews.

✨ Features

Create and fetch users

Store date of birth (DOB) in database

Calculate age dynamically (not stored)

SQL-safe database access using SQLC

Input validation with go-playground/validator

Structured logging using Uber Zap

Middleware for:

Request ID injection

Request duration logging

Clean HTTP status codes & error handling

Pagination-ready API design

🛠 Tech Stack

Go

GoFiber

PostgreSQL

SQLC

Uber Zap

go-playground/validator

🚀 Getting Started
Prerequisites

Go 1.21+

PostgreSQL

SQLC

Install SQLC (Windows Recommended)

Download prebuilt binary:
https://github.com/sqlc-dev/sqlc/releases

🗄 Database Setup
1️⃣ Create Database
CREATE DATABASE users;

2️⃣ Run Migration
psql -U USER -d users -f db/migrations/001_create_users.sql

⚙ Generate SQLC Code
sqlc generate


This generates type-safe DB access code in:

/db/sqlc/

▶ Run the Server
go mod tidy
go run cmd/server/main.go


Server will start on:

http://localhost:3000

📌 API Endpoints
➤ Create User

POST /users

{
  "name": "Anshu Singh",
  "dob": "1999-12-30"
}


Response

{
  "id": 1,
  "name": "Anshu Singh",
  "dob": "1999-12-30",
  "created_at": "2024-12-30T10:00:00Z"
}

➤ Get User by ID

GET /users/{id}

Response

{
  "id": 1,
  "name": "Anshu Singh",
  "dob": "1999-12-30",
  "age": 25
}

🧠 Design Decisions

Age is not stored → avoids stale data

SQLC ensures compile-time SQL safety

Service layer handles business logic

Middleware improves observability

Zap logger for structured logging

🧪 Testing

Age calculation is isolated and easily testable:

go test ./...

🐳 Optional (Docker Support)

Docker support can be added using Docker Compose for:

PostgreSQL

API Server

(Ask if you want this added)

🎯 Interview Ready Talking Points

Why SQLC over ORM

Why dynamic age calculation

Layered architecture benefits

Middleware for observability

Structured logging in production

📄 License

This project is open-source and available for learning and educational purposes.

🙌 Author

Anshu Singh
MCA Student | Aspiring Software Developer
