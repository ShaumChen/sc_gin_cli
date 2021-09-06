package kernel

import (
	"scgin/context"
	"scgin/exception"
)

var Middleware []context.HandlerFunc

func Load()  {
	Middleware = []context.HandlerFunc{
		exception.Exception,
	}
}