package main

import (
	"awesomeProject/Practice/ConnectLT"
	"awesomeProject/Practice/ControllerLT"
	"awesomeProject/Practice/middlewaresPrac"
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectLT.ConnectLt()
	r:=gin.Default()
	api := r.Group("/api",middlewaresPrac.AuthLt())
	{
		api.GET("/Select",ControllerLT.Select)
		api.POST("/InsertInto",ControllerLT.InsertInto)
		api.POST("/Login",ControllerLT.PrintLnToken)
	}
	r.GET("/Select",ControllerLT.Select)
	r.POST("/InsertInto",ControllerLT.InsertInto)
	r.POST("/Login",ControllerLT.PrintLnToken)
	r.Run()
}
