package ControllerB2

import (
	"awesomeProject/B2/ConnectLoginB2"
	"awesomeProject/B2/InforB2"
	"awesomeProject/B2/ServicerJwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Select(c *gin.Context)  {
	var arr []InforB2.Sinhviens
	ConnectLoginB2.ConnectB2()

	ConnectLoginB2.DB.Find(&arr)
	c.JSON(http.StatusOK,gin.H{"Data ":arr})
}

func Insert(c *gin.Context)  {
	ConnectLoginB2.ConnectB2()
	var Input InforB2.Sinhviens

	if err:= c.ShouldBindJSON(&Input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error ":err.Error()})
		return
	}
	hash,_ := HashPassword(Input.Passwordsv)
	Admin := InforB2.Sinhviens{
		Masv:       Input.Masv,
		Hotensv:    Input.Hotensv,
		Passwordsv: hash,
		Gmail:      Input.Gmail,
		Sodtsv:     Input.Sodtsv,
		Diachi:     Input.Diachi,
	}
	ConnectLoginB2.DB.Create(&Admin)
	c.JSON(http.StatusOK,gin.H{"Data ":Admin})
}

func Login(c *gin.Context) {
	ConnectLoginB2.ConnectB2()
	var Input InforB2.Sinhviens

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin InforB2.Sinhviens
	result := ConnectLoginB2.DB.Where(&InforB2.Sinhviens{
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
	k:= ServicerJwt.NewJWTService().GenerateToken(Input.Masv,Input.Masv)

	c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})
	c.JSON(http.StatusOK, gin.H{"Login ": k})

}

