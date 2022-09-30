package main

import (
	"awesomeProject/Students/ControllerStudents"
	"github.com/gin-gonic/gin"
)

func main() {
	r:=gin.Default()
	r.GET("/Select",ControllerStudents.Select)
	r.Run()

}















