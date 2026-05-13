package controller

import (
	"net/http"
	"regexp"

	"github.com/bernaddwiki/backend/dto"
	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

func NewRegisterController() *RegisterController {
	return &RegisterController{}
}

func (l *RegisterController) Register(c *gin.Context) {
	var body dto.RegisterBody
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

	// var emailRgx, _ = regexp.MatchString(
	// 	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, body.Email)

	// if !emailRgx {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "Email is not valid",
	// 	})
	// 	return
	// }

	defer func() {
		recover()
	}()
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

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Register success",
		"data": body,
	})
}
