package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Dob  string `json:"dob" validate:"required"`
}
