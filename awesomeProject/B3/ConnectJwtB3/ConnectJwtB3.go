package ConnectJwtB3

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectB3()  {
	dsn := "host=localhost user=postgres password=default dbname=SinhViens port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connect fail")
	} else {
		fmt.Println("Connect successfully")
	}
	DB=db
}
