package feature

import "scps-backend/feature/auth/domain/entities"

type User struct {
	Id         string  `json:"id" bson:"id"`
	InsurdNbr  string  `json:"insurdNbr" bson:"insurdNbr"`
	Nbr        int     `json:"nbr" bson:"nbr"`
	Name       string  `json:"name" bson:"name"`
	Age        int     `json:"age" bson:"age"`
	Email      string  `json:"email" bson:"email"`
	Password   string  `json:"password" bson:"password"`
	Phone      string  `json:"phone" bson:"phone"`
	JobYears   int     `json:"jobYears" bson:"jobYears"`
	CurrentJob string  `json:"currentJob" bson:"currentJob"`
	Salary     float32 `json:"salary" bson:"salary"`
	Permission string  `json:"permission" bson:"permission"`
	Son        []Son   `json:"son,omitempty" bson:"son"`
}

type Son struct {
	InsurdNbr string `json:"insurdNbr" bson:"insurdNbr"`
	Nbr       int    `json:"nbr" bson:"nbr"`
	Status    string `json:"status" bson:"status"`
	Name      string `json:"name" bson:"name"`
}

type Account interface {
	User | entities.Login | entities.SetEmail |
		entities.ReciveOTP | entities.SetPwd
}
