package controller

import (
	"log"
	"net/http"
	"scps-backend/api/controller/model"
	"scps-backend/core"
	"scps-backend/feature"
	"scps-backend/feature/auth/domain/entities"
	"scps-backend/feature/auth/usecase"
	util "scps-backend/util/token"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase usecase.AuthUsecase
}

// HANDLE WITH LOGIN ACCOUNT REQUEST
func (ic *AuthController) LoginRequest(c *gin.Context) {
	log.Println("************************ RECIEVE LOGIN REQUEST ************************")
	var loginParms entities.Login
	if !core.IsDataRequestSupported(&loginParms, c) {
		return
	}
	authParams := &usecase.LoginParams{}
	authParams.Data = &loginParms
	resulat := ic.AuthUsecase.Login(c, authParams)
	if err := resulat.Err; err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	secret := core.RootServer.SECRET_KEY
	user := resulat.Data.(*feature.User)
	token, err := util.CreateAccessToken(user.Id, secret, 2, user.Permission)
	if err != nil {
		c.JSON(500, model.ErrorResponse{Message: err.Error()})
		return
	}
	// log.Printf("TOKEN %s", token)
	user.Password = ""
	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "USER LOGGING SUCCESSFULY",
		Data: model.LoginResponse{
			Token:    token,
			UserData: user,
		},
	})

}
