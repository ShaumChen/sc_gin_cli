package kernel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yogo/model"
	"yogo/yogo"
)

func init() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
	yogo.Db = db
	if err != nil {
		fmt.Println("Unable to connect to the database", err)
	}
	db.AutoMigrate(&model.User{})

}
