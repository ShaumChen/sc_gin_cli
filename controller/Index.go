package controller

import (
	"scgin/context"
	"scgin/response"
)

func Index(context *context.Context) *response.Response {
	//panic("something error")
	return response.Resp().String(context.Context.FullPath())
}
