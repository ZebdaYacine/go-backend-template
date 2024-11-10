package repository

import (
	"context"
	"scps-backend/feature"
	"scps-backend/pkg/database"
)

type profileRepository struct {
	database database.Database
}

type ProfileRepository interface {
	GetProfile(c context.Context, userId string) (*feature.User, error)
}

func NewProfileRepository(db database.Database) ProfileRepository {
	return &profileRepository{
		database: db,
	}
}

func (r *profileRepository) GetProfile(c context.Context, userId string) (*feature.User, error) {
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
