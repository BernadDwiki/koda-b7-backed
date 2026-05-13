package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	// routing
	// register
	RegisterRouter(app)
	// login
	LoginRouter(app)
}
