package main

import (
	"awesomeProject/B1/Function"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.GET("/Select",Function.Select)
	r.POST("/Login",Function.Login)
	r.POST("/Login/Encode",Function.LoginEncodePass)
	r.POST("/Insert",Function.Insert)
	r.POST("/Insert/Encode",Function.InsertEncodePass)
	r.Run()

}
