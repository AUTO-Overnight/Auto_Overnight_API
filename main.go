package main

import (
	"auto_overnight_api/route"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.POST("/login", route.Login)
	r.POST("/send/stayout", route.SendStayOut)
	r.POST("/find/stayout", route.FindStayOutList)
	r.POST("/find/point", route.FindPointList)

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}
