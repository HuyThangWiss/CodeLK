package Infor

type Peoples struct {
	Peopleuser string `json:"peopleuser" gorm : "primaryKey" binding:"required"`
	Peoplepassword string `json:"peoplepassword" binding:"required"`
}
