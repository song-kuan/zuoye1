package controller

import(
    "seckill/service"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "strconv"
    "seckill/model"
    "seckill/rabbitmq"
    "encoding/json"
)

type ProductController struct{
    Ctx            iris.Context
    RabbitMQ       *rabbitmq.RabbitMQ
    ProductService service.ProductService
}

//获取商品信息
//http://localhost:8080/product/info
func (pc *ProductController)GetInfo()mvc.View{
    ProductArray,_:=pc.ProductService.GetAllProduct()
    if len(ProductArray)==0{
		 return mvc.View{
            Name:"error/not_found.html",
            Data:iris.Map{
                "msg":"商品不存在",
            },
        }
    }
    for i:=range ProductArray{
        ProductArray[i].Productnum-=1
    }
    return mvc.View{
        Name:"product/find_all_product.html",
        Data:iris.Map{
            "ProductArray":ProductArray,
        },
    }
}

//抢购指定商品
//http://localhost:8080/product/order/{id}
func (pc *ProductController)GetOrderBy(id int64)mvc.View{
    product,err:=pc.ProductService.GetProductById(id)
    if err!=nil{
        pc.Ctx.Application().Logger().Debug(err)
    }
    if product.Id==0{
        return mvc.View{
            Name:"error/not_found.html",
            Data:iris.Map{
                "msg":"指定商品不存在",
            },
        }
    }
    //生成订单
    userString:=pc.Ctx.GetCookie("uid")
    userID,_:=strconv.Atoi(userString)
    //生成消息体，插入rabbitmq
    message:=model.NewMessage(int64(userID),int64(id),0)
    byteMessage,err:=json.Marshal(message)
    err=pc.RabbitMQ.PublishSimple(string(byteMessage))
    
    if err!=nil{
        return mvc.View{
            Name:"product/product_order.html",
            Data:iris.Map{
                "msg":"生成订单失败",  
            },
        }
    }
    
    return mvc.View{
        Name:"product/product_order.html",
        Data:iris.Map{
            "msg":"订单已提交",
        },
    }
}