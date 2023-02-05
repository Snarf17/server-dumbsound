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
	dsn := "root:VGrKSXOUC2NanLt0HoTa@tcp(containers-us-west-118.railway.app:7682)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected To Database")
}
