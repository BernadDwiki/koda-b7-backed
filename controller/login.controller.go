package controller

import (
	"net/http"
	"regexp"

	"github.com/bernaddwiki/backend/dto"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (r *LoginController) Login(c *gin.Context) {
	var body dto.LoginBody
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

	var emailRegex = regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

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
}
