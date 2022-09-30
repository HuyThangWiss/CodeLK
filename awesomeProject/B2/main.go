package main

import (
	"awesomeProject/B2/ControllerB2"
	"awesomeProject/B2/ServicerJwt"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()

	api := r.Group("/api",ServicerJwt.AuthorizeJWT())
	{
		api.GET("/Select",ControllerB2.Select)
		api.POST("/Post",ControllerB2.Insert)
	}
	r.POST("/login2",ControllerB2.Login)
	r.Run()
}
