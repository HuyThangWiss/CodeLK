package ConnectLT

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectLt()  {
	dsn := "host=localhost user=postgres password=default dbname=User port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connect fail")
	} else {
		fmt.Println("Connect successfully")
	}
	DB=db
}
