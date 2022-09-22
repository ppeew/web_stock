package controller

import (
	"net/http"
	"web_stock/data"
	"web_stock/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
func RegisterPost(c *gin.Context) {
	if c.PostForm("password1") != c.PostForm("password2") {
		c.String(200, "密码不一致，请重新输入")
		return
	}
	u := model.User{
		Name:   c.PostForm("username"),
		Passwd: c.PostForm("password1"),
		Email:  c.PostForm("email"),
	}
	err := model.DB.Where("name=?", u.Name).First(&model.User{}).Error
	if err == gorm.ErrRecordNotFound {
		model.DB.Create(&u)
	} else if err != nil {
		c.String(200, "系统错误,请检查是否输入了中文字符")
		return
	} else {
		c.String(200, "创建用户失败,已经存在该用户")
		return
	}
	c.String(200, "创建用户成功")
}

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func LoginPost(c *gin.Context) {
	name, passwd := c.PostForm("username"), c.PostForm("password")
	if model.DB.Where("name=? and passwd=?", name, passwd).First(&model.User{}).Error != gorm.ErrRecordNotFound {
		c.String(http.StatusOK, "登录成功")
	} else {
		c.String(http.StatusOK, "账号或密码错误")
	}
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
