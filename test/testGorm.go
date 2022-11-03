package main

import (
	"GinChat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:13306)/ginchat?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}

	db.AutoMigrate(&models.UserBasic{})
	//user1 := &models.UserBasic{}
	//user1.Name = "crf"
	//db.Create(user1)

}
