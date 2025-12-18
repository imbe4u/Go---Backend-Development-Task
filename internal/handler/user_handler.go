package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"user-api/internal/models"
	"user-api/internal/repository"
	"user-api/internal/service"
)

type UserHandler struct {
	Repo     *repository.UserRepository
	Log      *zap.Logger
	Validate *validator.Validate
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return c.Status(400).JSON("DOB must be YYYY-MM-DD")
	}

	user, err := h.Repo.Create(c.Context(), req.Name, dob)
	if err != nil {
		h.Log.Error("create user failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Repo.GetByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  service.CalculateAge(user.Dob),
	})
}
