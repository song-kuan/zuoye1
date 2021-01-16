//rabbitMQ消息体
package model

//简单的消息体
type Message struct {
	Productid    int64
	Userid       int64
	Orderstatus  int64
}

//创建结构体
func NewMessage(userId int64, productId int64,orderStatus int64) *Message {
	return &Message{Userid: userId, Productid: productId,Orderstatus: orderStatus}
}