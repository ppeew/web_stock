package controller

import (
	"net/http"
	"web_stock/data"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Data(c *gin.Context) {
	var str string
	if data.ISOK {
		str = "可以"
	} else {
		str = "不适合"
	}
	c.HTML(http.StatusOK, "data.html", str)
}
