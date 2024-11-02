package controller

import (
	"log"
	"net/http"
	"scps-backend/api/controller/model"

	"github.com/gin-gonic/gin"
)

type PensionController struct {
}

// HANDLE WITH LOGIN ACCOUNT REQUEST
func (ic *PensionController) PensionRequest(c *gin.Context) {
	log.Println("************************ RECIEVE PENSION REQUEST ************************")

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "SCPC API PING SUCCESSFUL",
		Data:    "HEY HOW CAN I HELP YOU SIR?",
	})

}
