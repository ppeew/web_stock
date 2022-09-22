package main

import (
	"web_stock/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// data.GetData()
	// 服务器
	e := gin.Default()
	router.Router(e)
	e.Run(":8080")
}
