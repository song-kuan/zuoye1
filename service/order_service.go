package service

import(
    "seckill/model" 
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
	"log"
)

type OrderService interface{
    GetOrderById(int64)(model.Order,error)
    DeleteOrderById(int64)(model.Order,error)
    GetAllOrder(int64)([]*model.Order,error)
    //rabbitMQ插入订单
    InsertOrderByMessage(*model.Message)(model.Order,error)
}

type orderService struct{
    DB *xorm.Engine 
}

func NewOrderService(db *xorm.Engine)OrderService{
    return &orderService{DB:db}
}

//根据ID获取订单
func (os *orderService)GetOrderById(ID int64)(model.Order,error){
    order:=model.Order{}
    db:=os.DB
    _,err:=db.Table("order").Where("id=?",ID).Get(&order)
    if err!=nil{
        log.Println(err)
    }
    return order,err
}

//根据ID删除订单
func (os *orderService)DeleteOrderById(ID int64)(model.Order,error){
    order:=model.Order{}
    db:=os.DB
    _,err:=db.Table("order").Where("id=?",ID).Delete(&order)
    if err!=nil{
        log.Println(err)
    }
    return order,err
}
//获取指定用户的全部订单
func (os *orderService)GetAllOrder(userID int64)([]*model.Order,error){
    var orders []*model.Order
    db:=os.DB
    err:=db.Table("order").Where("userid=?",userID).Find(&orders)
    if err!=nil{
        log.Println(err)
    }
    return orders,err
}

//从rabbitMQ插入订单
func (os *orderService)InsertOrderByMessage(message *model.Message)(model.Order,error){
    order:=model.Order{
        Userid: message.Userid,
        Productid: message.Productid,
        //订单状态:待处理
        Orderstatus: message.Orderstatus,
    }
    db:=os.DB
    _,err:=db.Table("order").Insert(&order)
    if err!=nil{
        log.Println(err)
    }
    return order,err
}
