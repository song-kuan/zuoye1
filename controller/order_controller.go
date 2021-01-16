package controller

import(
    "seckill/service"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
    "strconv"
)

type OrderController struct{
    Ctx iris.Context
    OrderService service.OrderService
}

//http://localhost:8080/order/info
func (oc *OrderController)GetInfo()mvc.View{
    userString:=oc.Ctx.GetCookie("uid")
    userID,_:=strconv.Atoi(userString)
    orderArray,_:=oc.OrderService.GetAllOrder(int64(userID))
    if len(orderArray)==0{
        return mvc.View{
            Name:"error/not_found.html",
            Data:iris.Map{
                "msg":"用户订单为空",
            },
        }
    }
    return mvc.View{
        Name:"order/get_all_order.html",
        Data:iris.Map{
            "orderArray":orderArray,
        },
    }
}

//http://localhost:8080/order/find/{id}
/*func (oc *OrderController)GetFindBy(id int64)mvc.View{
    order,_:=oc.OrderService.GetOrderById(id)
    if order.Id==0{
        return mvc.View{
            Name:"error/not_found.html",
            Data:iris.Map{
                "msg":"指定订单不存在",
            },
        }
    }
    return mvc.View{
        Name:"order/get_order_by_id.html",
        Data:iris.Map{
            "orderID":  order.Id,
            "userID":   order.Userid,
            "productID": order.Productid,
            "status": order.Orderstatus,
        },
    }
}*/

//http://localhost:8080/order/erase/{id}
func (oc *OrderController)GetEraseBy(id int64)mvc.View{
    oc.OrderService.DeleteOrderById(id)
    return mvc.View{
        Name:"order/erase_order.html",
        Data:iris.Map{
            "msg":"订单删除成功",
        },
    }
}
