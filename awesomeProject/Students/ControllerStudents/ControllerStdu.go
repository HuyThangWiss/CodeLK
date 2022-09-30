package ControllerStudents

import (
	"awesomeProject/Students/ConnectStudents"
	"awesomeProject/Students/InforStudents"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Users struct {
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Grade     string `form:"grade"`
	Gender    string `form:"gender"`
	Gmail     string `form:"gmail"`
}

func Select(c *gin.Context){
	ConnectStudents.ConnectBStudent()
	var arr[] InforStudents.Students
	var user Users
	err:=c.ShouldBindQuery(&user)

	if err!= nil{
		c.JSON(http.StatusOK,gin.H{"Error ":err.Error()})
		return
	}
	result := ConnectStudents.DB.Where(InforStudents.Students{
		Firstname:  user.Firstname,
		Lastname:  user.Lastname,
		Grade:     user.Grade,
		Gender:    user.Gender,
		Gmail:     user.Gmail,
	}).Find(&arr)
	if result.Error != nil{
		c.JSON(http.StatusOK,gin.H{"Error ":err.Error()})
		return
	}
	if len(arr) == 0{
		c.JSON(http.StatusOK,gin.H{"Empty":"Not exists Data"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"Data ":arr})
	return
}
/*
func InsertInto(c *gin.Context)  {


	ConnectStudents.ConnectBStudent()

	var Input InforStudents.Students

	if err:=c.ShouldBind(&Input);err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if Input.Gmail =="" || Input.Id == ""|| Input.Firstname ==""|| Input.Lastname == "" || Input.Gender == ""|| Input.Grade ==""{
		c.JSON(http.StatusOK,gin.H{"Error":"Check Emplty"})
		return
	}

	Admin :=InforStudents.Students{
		Id:        Input.Id,
		Firstname: Input.Firstname,
		Lastname:  Input.Lastname,
		Grade:     Input.Grade,
		Gender:    Input.Gender,
		Gmail:     Input.Gmail,
	}
	ConnectStudents.DB.Create(&Admin)
	return

}*/













