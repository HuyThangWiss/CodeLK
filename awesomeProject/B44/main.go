package main

import (
	"awesomeProject/B44/Controler"
	"github.com/gin-gonic/gin"

)

//http://localhost:8080/Select

func main() {

	router := gin.Default()
	router.GET("/task",Controler.B55)
	router.Run(":8081")
}
