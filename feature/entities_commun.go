package feature

import "scps-backend/feature/auth/domain/entities"

type User struct {
	Id         string  `json:"id" bson:"id"`
	Name       string  `json:"name" bson:"name"`
	Age        int     `json:"age" bson:"age"`
	Email      string  `json:"email" bson:"email"`
	Password   string  `json:"password" bson:"password"`
	Phone      string  `json:"phone" bson:"phone"`
	JobYears   int     `json:"jobYears" bson:"jobYears"`
	CurrentJob string  `json:"currentJob" bson:"currentJob"`
	Salary     float32 `json:"salary" bson:"salary"`
	Permission string  `json:"permission" bson:"permission"`
}

type Account interface {
	User | entities.Login
}
