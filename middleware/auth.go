package middleware

import "github.com/kataras/iris/v12"

func Auth(ctx iris.Context){
	uid:=ctx.GetCookie("uid")
	if uid==""{
		//ctx.Application().Logger().Debug("必须先登录!")
		ctx.WriteString("<h1>必须先登录!</h1>")
		ctx.Redirect("/user/login")
		return
	}
	//ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}