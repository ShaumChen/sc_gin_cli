package controller

import (
	"golang.org/x/time/rate"
	"scgin/component/limiter"
	"scgin/component/lock"
	"scgin/context"
	"scgin/kernel"
	"scgin/model"
	"scgin/response"
	"strconv"
	"time"
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

func TestLimiter(context *context.Context) *response.Response {
	l := limiter.NewLimiter(rate.Every(1*time.Second), 1, context.ClientIP())
	if !l.Allow() {
		return response.Resp().String("error")
	}
	return response.Resp().String("success")
}

func TestLock(context *context.Context) *response.Response {
	l := lock.NewLock("test", 10*time.Second)
	defer l.Release()
	if l.Get() {
		return response.Resp().String("拿锁成功")
	}
	return response.Resp().String("拿锁失败")
}

func TestBlock(context *context.Context) *response.Response {
	l := lock.NewLock("test", 10*time.Second)
	defer l.Release()
	if l.Block(5 * time.Second) {
		return response.Resp().String("拿锁成功")
	}
	return response.Resp().String("拿锁失败")
}

func TestCreateUser(context *context.Context) *response.Response {
	//user := &model.User{
	//	UserName: "test",
	//	Password: "123456",
	//}
	//_, err := (*sql.DB).Exec("insert into user (user_name, password) values (?, ?)","sdhjds","dsadas")
	//if err != nil {
	//	return nil
	//}
	row := kernel.Db.First(&model.User{}).Value
	return response.Resp().Json(row)
}
