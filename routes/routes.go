package routes

import (
	"github.com/gin-gonic/gin"
	"scgin/controller"
	"scgin/kernel"
	"scgin/middleware"
)

func config(router group) {
	router.Group("/api", func(api group) {
		api.Group("/user", func(user group) {
			user.Registered(GET, "/info", controller.Index, middleware.M3)
			user.Registered(GET, "/order", controller.Index)
			user.Registered(GET, "/money", controller.Index)
		}, middleware.M2)
	}, middleware.M1)
}

func Load(r *gin.Engine) {
	router := newRouter(r)
	router.Group("", func(g group) {
		config(g)
	}, kernel.Middleware...)
	r.GET("/", convert(controller.Index))
}
