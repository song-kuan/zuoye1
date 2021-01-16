package model

type Product struct {
	Id           int64  `xorm:"pk autoincr" json:"id"`
	Productname  string `json:"productname"`
	Productnum   int64  `json:"productnum"`
}