package main

import (
	"github.com/gin-gonic/gin"
	"scgin/kernel"
	"scgin/routes"
)

func main() {
	r := gin.Default()
	kernel.Load()
	routes.Load(r)
	r.Run("0.0.0.0:8888")
}
