package public

import (
	"scps-backend/api/controller"
	"scps-backend/core"
	"scps-backend/feature/auth/domain/repository"
	"scps-backend/feature/auth/usecase"
	"scps-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(db database.Database, group *gin.RouterGroup) {
	ir := repository.NewAuthRepository(db)
	uc := usecase.NewAuthUsecase(ir, core.USER)
	ic := &controller.AuthController{
		AuthUsecase: uc, // usecase for insured operations
	}
	group.GET("login", ic.LoginRequest)
}
