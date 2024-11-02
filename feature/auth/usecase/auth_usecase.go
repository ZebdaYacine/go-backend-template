package usecase

import (
	"context"
	"fmt"
	"scps-backend/feature/auth/domain/entities"
	"scps-backend/feature/auth/domain/repository"
)

type LoginParams struct {
	Data *entities.Login
}

type LoginResult struct {
	Data any
	Err  error
}

func (l *LoginParams) validator() error {
	loginParams := l.Data
	if loginParams.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if loginParams.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	return nil
}

type AuthUsecase interface {
	//AUTH FUNCTIONS
	Login(c context.Context, data *LoginParams) *LoginResult
}

type authUsecase struct {
	repo       repository.AuthRepository
	collection string
}

func NewAuthUsecase(repo repository.AuthRepository, collection string) AuthUsecase {
	return &authUsecase{
		repo:       repo,
		collection: collection,
	}
}

// Login implements UserUsecase.
func (u *authUsecase) Login(c context.Context, data *LoginParams) *LoginResult {
	if err := data.validator(); err != nil {
		return &LoginResult{Err: err}
	}
	loginResult, err := u.repo.Login(c, data.Data)
	if err != nil {
		return &LoginResult{Err: err}
	}
	return &LoginResult{Data: loginResult}
}
