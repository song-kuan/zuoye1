package router

import(
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "seckill/controller"
    "seckill/datasource"
    "seckill/service"
    "github.com/kataras/iris/v12/sessions"
    "time"
    "seckill/middleware"
    "seckill/rabbitmq"
)

func NewHandle(app *iris.Application){
    //设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")
	
	engine:=datasource.NewEngine()
	//简单队列模式
	rabbitmq:=rabbitmq.NewRabbitMQSimple("Simple")
	
	//http://localhost:8080/order
	orderService:=service.NewOrderService(engine)
	orderParty:=app.Party("/order")
	order:=mvc.New(orderParty)
	orderParty.Use(middleware.Auth)
	order.Register(
	    orderService,
	)
	order.Handle(new(controller.OrderController))
	
	//http://localhost:8080/user
	userService:=service.NewUserService(engine)
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	userParty:=app.Party("/user")
	user:=mvc.New(userParty)
	user.Register(
	   userService,
	   sessManager.Start,
	)
	user.Handle(new(controller.UserController))
	
	//http://localhost:8080/product
	productService:=service.NewProductService(engine)
	productParty:=app.Party("/product")
	product:=mvc.New(productParty)
	productParty.Use(middleware.Auth)
	product.Register(
	    rabbitmq,
	    productService,
	)
	product.Handle(new(controller.ProductController))
	
	//订单处理
	go rabbitmq.ConsumeSimple(orderService,productService)  
}