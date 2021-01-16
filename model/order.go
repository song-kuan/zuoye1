package model

type Order struct {
    Id          int64 `xorm:"pk autoincr" json:"id"`
	Userid      int64 `json:"userid"`
	Productid   int64 `json:"productid"`
	Orderstatus int64 `json:"orderstatus"`
}

//订单状态,等待付款,成功状态,失败
const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)
