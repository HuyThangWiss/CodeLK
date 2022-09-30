package main

import (
	"awesomeProject/B3/ConnectJwtB3"
	"awesomeProject/B3/ControllerJwtB3"
	"awesomeProject/B3/MiddwareJwtB3"

	//"awesomeProject/B3/MiddwareJwtB3"
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectJwtB3.ConnectB3()

	r:= gin.Default()
	api := r.Group("/api",MiddwareJwtB3.AuthorizeJWT())
	{
		api.GET("/Select",ControllerJwtB3.Select)
		api.POST("/Insert",ControllerJwtB3.Insert)
	}
	r.POST("/Login",ControllerJwtB3.Login)
	r.GET("/Select",ControllerJwtB3.Select)
	r.POST("/token",ControllerJwtB3.LoginToken)
	r.POST("/search/:id=?")
	r.Run()

}
