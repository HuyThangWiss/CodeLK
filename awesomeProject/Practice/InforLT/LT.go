package InforLT

type Users struct {
	Name string `json:"name" binding:"required" gorm:"primaryKey" gorm:"index"`
	Username string `json:"username" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Id int `json:"id"  gorm:"AUTO_INCREMENT"`

}

type SignUpInput struct {
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`