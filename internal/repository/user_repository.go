package repository

import (
	"context"
	"time"

	"user-api/db/sqlc"
)

type UserRepository struct {
	Q *sqlc.Queries
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
	return r.Q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	return r.Q.GetUserByID(ctx, id)
}
