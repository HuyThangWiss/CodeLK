package ControllerLT

import (
	"awesomeProject/Practice/AuthPrac"
	"awesomeProject/Practice/ConnectLT"
	"awesomeProject/Practice/InforLT"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Body struct {
	ID int `json:"Id" orm:="auto"`
	Data interface{}

}

func Select(c *gin.Context)  {
	ConnectLT.ConnectLt()

	brr := make(map[int]Body)
	var arr []InforLT.Users

	ConnectLT.DB.Find(&arr)

	for k,i:= range arr {
		brr[k+1]=Body{
			ID:   k+1,
			Data: i,
		}
	c.JSON(http.StatusOK,gin.H{"Daata":brr})
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func InsertInto(c *gin.Context)  {
	ConnectLT.ConnectLt()
	var Input InforLT.Users

	if err:= c.ShouldBindJSON(&Input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error ":err.Error()})
		return
	}
	hash,_ := HashPassword(Input.Password)
	Admin := InforLT.Users{

		Name:     Input.Name,
		Username: Input.Username,
		Email:    Input.Email,
		Password: hash,
	}
	err:=ConnectLT.DB.Create(&Admin)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"Errror":err.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"Data ":Admin})
	return
}

type TokenRequest struct {
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func PrintLnToken(c *gin.Context)  {
	ConnectLT.ConnectLt()
	var Input InforLT.Users

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin InforLT.Users
	result := ConnectLT.DB.Where(&InforLT.Users{
		Email:       Input.Email,

	}).First(&admin)

	if result.Error != nil {
		c.JSON(http.StatusOK,gin.H{"errpr":"User not ecits"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(Input.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}
	tokenString, err:= AuthPrac.GenerateJWT(Input.Email, Input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login sucessful ": tokenString})
}


//func Login(c *gin.Context) {
//	ConnectLT.ConnectLt()
//	var Input InforLT.Users
//
//	if err := c.ShouldBindJSON(&Input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
//		return
//	}
//	var admin InforLT.Users
//	result := ConnectLT.DB.Where(&InforLT.Users{
//		Email: Input.Email,
//	}).First(&admin)
//
//	if result.Error != nil {
//		c.JSON(http.StatusOK,gin.H{"errpr":"User not ecits"})
//		return
//	}
//	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(Input.Password))
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})
//
//}

//func Select(c *gin.Context)  {
//	ConnectLT.ConnectLt()
//
//	brr := make(map[int]InforLT.Users)
//	var arr [] InforLT.Users
//
//	ConnectLT.DB.Find(&arr)
//	for k,i:= range arr {
//		brr[k+1]=i
//	}
//
//	c.JSON(http.StatusOK,gin.H{"Data ":brr})
//}