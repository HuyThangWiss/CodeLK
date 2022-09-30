package main

import (
	"awesomeProject/B3/ConnectJwtB3"
	"awesomeProject/B3/InforJwtB3"
	"encoding/json"
	"fmt"
)

func main() {
	var arr []InforJwtB3.Sinhviens
	ConnectJwtB3.ConnectB3()
	ConnectJwtB3.DB.Find(&arr)
	byteArray, err := json.MarshalIndent(arr, "", "  ")
	if err != nil{
		return
	}
	fmt.Println(string(byteArray))
}
