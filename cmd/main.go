package main

import (
	"github.com/bernaddwiki/backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	router.InitRouter(route)
	route.Run("localhost:8080")
}
