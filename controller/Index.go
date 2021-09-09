package controller

import (
	"scgin/context"
	"scgin/response"
	"strconv"
)

func Index(context *context.Context) *response.Response {
	//panic("something error")
	return response.Resp().String(context.Context.FullPath())
}

func TestSetSession(context *context.Context) *response.Response {
	context.Session().Set("msg", "PHPer")
	return response.Resp().String("set session")
}

func TestGetSession(context *context.Context) *response.Response {
	context.Session().Get("msg")
	return response.Resp().String("get session")
}

func TestRemoveSession(context *context.Context) *response.Response {
	context.Session().Remove("msg")
	return response.Resp().String("remove session")
}

func TestCoroutineSetSession(context *context.Context) *response.Response {
	session := context.Session()
	for i := 0; i < 100; i++ {
		go func(index int) {
			session.Set("msg"+strconv.Itoa(index), index)
		}(i)
	}
	return response.Resp().String("coroutine set session")
}
