package router

import (
	"github.com/bernaddwiki/backend/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(app *gin.Engine) {
	loginController := controller.NewLoginController()
	app.POST("/login", loginController.Login)
}
