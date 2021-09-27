package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yogo/kernel"
	"yogo/model"
	"yogo/routes"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("Unable to connect to the database", err)
	}
	db.AutoMigrate(&model.User{})
	kernel.Db = db
	r := gin.Default()
	kernel.Load()
	routes.Load(r)
	r.Run("0.0.0.0:8888")
}
