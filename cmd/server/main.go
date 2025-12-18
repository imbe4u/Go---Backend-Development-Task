package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/repository"
	"user-api/internal/routes"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost:5432/users?sslmode=disable")

	q := sqlc.New(db)
	repo := &repository.UserRepository{Q: q}

	logg := logger.New()
	validate := validator.New()

	h := &handler.UserHandler{
		Repo: repo,
		Log:  logg,
		Validate: validate,
	}

	app := fiber.New()
	app.Use(middleware.RequestLogger(logg))

	routes.Register(app, h)

	log.Fatal(app.Listen(":3000"))
}
