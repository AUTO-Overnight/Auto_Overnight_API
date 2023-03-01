package main

import (
	"auto_overnight_api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", route.Login)
	r.POST("/send/stayout", route.SendStayOut)
	r.POST("/find/stayout", route.FindStayOutList)
	r.POST("/find/point", route.FindPointList)

	r.Run(":8081")
}
