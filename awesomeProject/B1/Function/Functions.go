package Function

import (
	"awesomeProject/B1/Connect"
	"awesomeProject/B1/Infor"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Select(c *gin.Context)  {
	var Input []Infor.Peoples
	Connect.Connect()
	Connect.DB.Find(&Input)
	c.JSON(http.StatusOK,gin.H{"Data ": Input})

}
func Insert(c *gin.Context)  {
	Connect.Connect()
	var Input Infor.Peoples

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	people := Infor.Peoples{
		Peopleuser:     Input.Peopleuser,
		Peoplepassword: Input.Peoplepassword,
	}
	Connect.DB.Create(&people)

	c.JSON(http.StatusOK,gin.H{"Data ": people})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func InsertEncodePass(c *gin.Context)  {
	Connect.Connect()
	var Input Infor.Peoples

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	hash,_ := HashPassword(Input.Peoplepassword)
	people := Infor.Peoples{
		Peopleuser:     Input.Peopleuser,
		Peoplepassword: hash,
	}
	Connect.DB.Create(&people)

	c.JSON(http.StatusOK,gin.H{"Data ": people})

}

func LoginEncodePass(c *gin.Context)  {
	Connect.Connect()
	var Input Infor.Peoples

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	var a Infor.Peoples
	hash,_ := HashPassword(Input.Peoplepassword)
	if Input.Peopleuser == "" || Input.Peoplepassword == ""{
		c.JSON(http.StatusOK,gin.H{"Login ": "Input Empty"})
	}else if err2 := Connect.DB.Where("Peopleuser = ?",Input.Peopleuser).First(&a).Error;err2 == nil{
		if err3:= Connect.DB.Where("Peoplepassword = ?",hash).First(&a).Error;err3 == nil{
			c.JSON(http.StatusOK,gin.H{"Login ": "Successfull"})
		}else{
			c.JSON(http.StatusOK,gin.H{"Login ": "Check password"})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{"Login ": "Check User"})
	}
}

func Login(c *gin.Context)  {
	Connect.Connect()
	var Input Infor.Peoples

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	var a Infor.Peoples
	if Input.Peopleuser == "" || Input.Peoplepassword == ""{
		c.JSON(http.StatusOK,gin.H{"Login ": "Input Empty"})
	}else if err2 := Connect.DB.Where("Peopleuser = ?",Input.Peopleuser).First(&a).Error;err2 == nil{
		if err3:= Connect.DB.Where("Peoplepassword = ?",Input.Peoplepassword).First(&a).Error;err3 == nil{
			c.JSON(http.StatusOK,gin.H{"Login ": "Successfull"})
		}else{
			c.JSON(http.StatusOK,gin.H{"Login ": "Check password"})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{"Login ": "Check User"})
	}
}



