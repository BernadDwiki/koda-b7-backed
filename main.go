package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

type LoginBody struct {
	Email    string `json:"email`
	Password string `json:"password"`
}

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func main() {
	// inisialisasi
	app := gin.Default()
	// routing
	// register
	app.POST("/register", func(c *gin.Context) {
		var body RegisterBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}

		if body.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Email is required",
			})
			return
		}

		if !emailRegex.MatchString(body.Email) {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Email is not valid",
			})
			return
		}

		if body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Password is required",
			})
			return
		}

		if len(body.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Password must be at least 6 characters",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "Register success",
			"data": body,
		})
	})
	// login
	app.POST("/login", func(c *gin.Context) {
		var body LoginBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal Server Error",
			})
			return
		}

		if body.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Email is required",
			})
			return
		}

		if !emailRegex.MatchString(body.Email) {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Email is not valid",
			})
			return
		}

		if body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Password is required",
			})
			return
		}

		if len(body.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Password must be at least 6 characters",
			})
			return
		}

		// dummy login
		if body.Email != "bernad@mail.com" || body.Password != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Invalid email or password",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "Login success",
		})
	})
	app.Run("localhost:8080")
}
