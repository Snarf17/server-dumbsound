package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// dsn := "root:@tcp(127.0.0.1:3306)/dumbsound?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:qDyWAxa56z04OgybAZaN@containers-us-west-189.railway.app:6092/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected To Database")
}
