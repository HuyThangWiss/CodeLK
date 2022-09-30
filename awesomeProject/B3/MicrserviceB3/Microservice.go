package MicrserviceB3

import (
	"awesomeProject/B3/ConnectJwtB3"
	"awesomeProject/B3/InforJwtB3"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Service(c *gin.Context)  {
	ConnectJwtB3.ConnectB3()
	posturl := "http://localhost:8080/api/List"
	r, err := http.NewRequest(http.MethodGet, posturl, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Auto", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtYXN2IjoiMDA1IiwicGFzc3dvcmRzdiI6IjEyMzQ1NjciLCJleHAiOjE2NjM5MzI2ODUsImlhdCI6MTY2MzY3MzQ4NX0.Zv9H-3dmMdBBgDjZ23dbeXbPOT6yg2kT1rZS1yg_PHk")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(err)
	}
	var tasks []InforJwtB3.Sinhviens
	json.Unmarshal(body, &tasks)
	c.JSON(200, gin.H{
		"body:": tasks,
	})
	return
}
