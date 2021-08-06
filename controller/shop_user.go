package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog/service"
)

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	login, err := service.Login(username, password)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, login)

}