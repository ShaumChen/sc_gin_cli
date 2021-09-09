package kernel

import (
	"scgin/context"
	"scgin/exception"
	"scgin/middleware/session"
)

var Middleware []context.HandlerFunc

func Load() {
	Middleware = []context.HandlerFunc{
		exception.Exception,
		session.Session,
	}
}
