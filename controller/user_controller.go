package controller

import(
    "seckill/service"
    "github.com/kataras/iris/v12"
    "seckill/model"
    "github.com/kataras/iris/v12/mvc"
    "github.com/kataras/iris/v12/sessions"
    "strconv"
    "net/http"
)

type UserController struct{
    Ctx iris.Context
    UserService service.UserService
    Session *sessions.Session
}

//用户注册
func (uc *UserController)GetRegister(ctx iris.Context)mvc.View{
    return mvc.View{
        Name:"user/register.html",
    }
}
func (uc *UserController)PostRegister(ctx iris.Context)mvc.Response{
    var(
		userName=uc.Ctx.FormValue("userName")
		password=uc.Ctx.FormValue("password")
    )
    flag:=uc.UserService.IsUnique(userName)
    //注册用户名已存在
    if flag==true{
        return mvc.Response{
            Path:"/user/register",
        }
    }
    user:=model.User{
		Username:     userName,
		Password:     password,
	}

	_,err:=uc.UserService.AddUser(user)
	uc.Ctx.Application().Logger().Debug(err)
	if err!=nil{
		return mvc.Response{
		    Path:"/user/error",
		}
	}
	uc.Ctx.Redirect("/user/login")
	return mvc.Response{
	    Path:"/user/login",
	}
}

//用户登陆
func (uc *UserController)GetLogin(ctx iris.Context)mvc.View{
    return mvc.View{
        Name:"user/login.html",
    }
}
func (uc *UserController)PostLogin(ctx iris.Context)mvc.Response{
   var (
		userName=uc.Ctx.FormValue("userName")
		password=uc.Ctx.FormValue("password")
	)
	
	user,_:=uc.UserService.IsLogIn(userName, password)
	if user.Id==0 {
		return mvc.Response{
			Path: "/user/login",
		}
	}
    
    //设置cookie和session
    uc.Ctx.SetCookie(&http.Cookie{Name:"uid",Value:strconv.FormatInt(user.Id,10),Path:"/"})
	uc.Session.Set("userId",strconv.FormatInt(user.Id,10))
	
	return mvc.Response{
		Path: "/product/info",
	}
}

//用户退出
func (uc *UserController)GetLogout(ctx iris.Context)mvc.Response{
	//删除session，下次需要从新登录
	uc.Session.Delete("userId")
	uc.Ctx.RemoveCookie("uid")
	return mvc.Response{
		Path:"/user/login",
	}
}
