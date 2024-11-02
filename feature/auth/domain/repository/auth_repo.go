package repository

import (
	"context"
	"scps-backend/feature"
	domain "scps-backend/feature/auth/domain/entities"
	"scps-backend/pkg/database"
)

type authRepository struct {
	database database.Database
}

type AuthRepository interface {
	Login(c context.Context, data *domain.Login) (*feature.User, error)
}

func NewAuthRepository(db database.Database) AuthRepository {
	return &authRepository{
		database: db,
	}
}

func (r *authRepository) Login(c context.Context, data *domain.Login) (*feature.User, error) {
	return &feature.User{
		Id:         "4j239412nk92u411233mnd",
		Name:       "Admin",
		Email:      "z9PzA@example.com",
		Permission: "user",
	}, nil
}
