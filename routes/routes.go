package routes

import (
	"github.com/gin-gonic/gin"
	"yogo/controller"
	"yogo/kernel"
	"yogo/middleware"
)

func config(router group) {
	router.Registered(GET, "/", controller.Index)
	router.Registered(GET, "/testSetSession", controller.TestSetSession)
	router.Registered(GET, "/testGetSession", controller.TestGetSession)
	router.Registered(GET, "/testRemoveSession", controller.TestRemoveSession)
	router.Registered(GET, "/testCoroutineSetSession", controller.TestCoroutineSetSession)
	router.Registered(GET, "/testLimiter", controller.TestLimiter)
	router.Registered(GET, "/testCreateUser", controller.TestCreateUser)
	router.Registered(GET, "/testGetUser", controller.TestGetUser)
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
}
