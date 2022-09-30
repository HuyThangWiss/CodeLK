package ControllerJwtB3

import (
	"awesomeProject/B3/ConnectJwtB3"
	"awesomeProject/B3/InforJwtB3"
	"awesomeProject/B3/TokenJwtB3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Select(c *gin.Context)  {
	var arr []InforJwtB3.Sinhviens
	ConnectJwtB3.ConnectB3()

	ConnectJwtB3.DB.Find(&arr)
	c.JSON(http.StatusOK,gin.H{"Data ":arr})
}

func Insert(c *gin.Context)  {
	ConnectJwtB3.ConnectB3()
	var Input InforJwtB3.Sinhviens

	if err:= c.ShouldBindJSON(&Input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error ":err.Error()})
		return
	}
	hash,_ := HashPassword(Input.Passwordsv)
	Admin := InforJwtB3.Sinhviens{
		Masv:       Input.Masv,
		Hotensv:    Input.Hotensv,
		Passwordsv: hash,
		Gmail:      Input.Gmail,
		Sodtsv:     Input.Sodtsv,
		Diachi:     Input.Diachi,
	}
	ConnectJwtB3.DB.Create(&Admin)
	c.JSON(http.StatusOK,gin.H{"Data ":Admin})
}

func Login(c *gin.Context) {
	ConnectJwtB3.ConnectB3()
	var Input InforJwtB3.Sinhviens

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin InforJwtB3.Sinhviens
	result := ConnectJwtB3.DB.Where(&InforJwtB3.Sinhviens{
		Masv:       Input.Masv,

	}).First(&admin)

	if result.Error != nil {
		c.JSON(http.StatusOK,gin.H{"errpr":"User not ecits"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Passwordsv), []byte(Input.Passwordsv))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})

}


func LoginToken(c *gin.Context) {
	ConnectJwtB3.ConnectB3()
	var Input InforJwtB3.Sinhviens

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin InforJwtB3.Sinhviens
	result := ConnectJwtB3.DB.Where(&InforJwtB3.Sinhviens{
		Masv:       Input.Masv,

	}).First(&admin)

	if result.Error != nil {
		c.JSON(http.StatusOK,gin.H{"errpr":"User not ecits"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Passwordsv), []byte(Input.Passwordsv))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}

	z:= TokenJwtB3.NewJWTServiceB3().GenerateTokenB3(Input.Masv,Input.Passwordsv)

	c.JSON(http.StatusOK, gin.H{"Login ": z})

}





