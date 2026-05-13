package router

import (
	"github.com/bernaddwiki/backend/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	registerController := controller.NewRegisterController()
	app.POST("/register", registerController.Register)
}
