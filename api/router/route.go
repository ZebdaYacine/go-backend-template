package router

import (
	"scps-backend/api/controller/middleware"
	"scps-backend/api/router/private"
	"scps-backend/api/router/public"

	"scps-backend/pkg"
	"scps-backend/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(db database.Database, gin *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Change to your Flutter web app's URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	gin.Use(cors.New(config))

	publicRouter := gin.Group("/")

	// All Public APIs
	public.NewPingRouter(db, publicRouter)
	public.NewLoginRouter(db, publicRouter)

	userRouter := gin.Group("/profile")

	//Middleware to verify AccessToken
	userRouter.Use(middleware.JwtAuthMiddleware(
		pkg.GET_ROOT_SERVER_SEETING().SECRET_KEY,
		"User"))

	//Auth API
	private.NewGetPensionRouter(db, userRouter)

	adminRouter := gin.Group("/admin")

	//Middleware to verify AccessToken
	adminRouter.Use(middleware.JwtAuthMiddleware(
		pkg.GET_ROOT_SERVER_SEETING().SECRET_KEY,
		"Admin"))

	private.NewUpdatePensionRateRouter(db, adminRouter)
}
