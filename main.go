package main

import(
    "github.com/kataras/iris/v12"
    "seckill/router"
    "seckill/config"
)

func main(){
    app:=iris.New()
    //注册views目录下的所有html文件
    app.RegisterView(iris.HTML("./views",".html"))
    //出错时访问页面
    app.OnAnyErrorCode(func(ctx iris.Context){
        ctx.View("error/error.html")
    })
    router.NewHandle(app)
    //默认地址为user/login
    app.Get("/",func(ctx iris.Context){
        ctx.View("user/login.html")
    })
    //获取config的端口号
    Config:=config.InitConfig()
    addr:=":"+Config.Port
    app.Run(
	    iris.Addr(addr),                              
		iris.WithoutServerError(iris.ErrServerClosed), 
		iris.WithOptimizations,                        
    )
}