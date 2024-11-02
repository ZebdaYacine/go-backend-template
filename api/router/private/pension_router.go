package private

import (
	"scps-backend/api/controller"
	"scps-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func NewGetPensionRouter(db database.Database, group *gin.RouterGroup) {
	// ir := repository.NewUserRepository(db)
	// uc := usecase.NewUserUsecase(ir, "")
	ic := &controller.PensionController{}
	group.GET("get-pension", ic.PensionRequest)
}

func NewUpdatePensionRateRouter(db database.Database, group *gin.RouterGroup) {
	// ir := repository.NewUserRepository(db)
	// uc := usecase.NewUserUsecase(ir, "")
	// ic := &controller.TestController{
	// 	UserUsecase: uc, // usecase for insured operations
	// }
	// group.GET("ping", ic.PingRequest)
}
