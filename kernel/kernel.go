package kernel

import (
	"yogo/context"
	"yogo/exception"
	"yogo/middleware/session"
)

var Middleware []context.HandlerFunc

func Load() {
	Middleware = []context.HandlerFunc{
		exception.Exception,
		session.Session,
	}
}
