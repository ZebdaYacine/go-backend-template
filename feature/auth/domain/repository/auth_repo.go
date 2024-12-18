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
	SetPassword(c context.Context, pwd string) (bool, error)
}

func NewAuthRepository(db database.Database) AuthRepository {
	return &authRepository{
		database: db,
	}
}

func (r *authRepository) Login(c context.Context, data *domain.Login) (*feature.User, error) {
	return &feature.User{
		Id:         "4j239412nk92u411233mnd",
		InsurdNbr:  "3423405",
		Nbr:        2,
		Name:       "zebda yacine",
		Email:      "zebdaadam1996@gmail.com",
		Password:   "admin",
		Permission: "User",
		Son: []feature.Son{
			{
				InsurdNbr: "3423403",
				Nbr:       2,
				Status:    "wife",
				Name:      "nasima",
			},
			{
				InsurdNbr: "3423409",
				Nbr:       1,
				Status:    "boy",
				Name:      "nasim",
			},
		},
	}, nil
}

func (r *authRepository) SetPassword(c context.Context, pwd string) (bool, error) {
	return true, nil
}
