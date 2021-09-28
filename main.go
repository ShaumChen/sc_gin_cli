package main

import (
	"github.com/gin-gonic/gin"
	"yogo/kernel"
	"yogo/routes"
	"yogo/yogo"
)

func main() {
	defer yogo.Db.Close()
	r := gin.Default()
	kernel.Load()
	routes.Load(r)
	r.Run("0.0.0.0:8888")
}
