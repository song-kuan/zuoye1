package service

import(
    "seckill/model" 
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
	"log"
	"errors"
)

type ProductService interface{
    GetProductById(int64)(model.Product,error)
    GetAllProduct()([]*model.Product,error)
    GetProductNumById(int64)(int64,error)
    UpdateServie(model.Product)(error)
    SubProductNum(productID int64)(error)
}

type productService struct{
    DB *xorm.Engine
}

func NewProductService(db *xorm.Engine)ProductService{
    return &productService{DB:db}
}

//根据ID寻找商品
func (ps *productService)GetProductById(ID int64)(model.Product,error){
    product:=model.Product{}
    db:=ps.DB
    _,err:=db.Table("product").Where("Id=?",ID).Get(&product)
    if err!=nil{
        log.Println(err)
    }
    return product,err
}

//获取全部商品信息
func (ps *productService)GetAllProduct()([]*model.Product,error){
    var products []*model.Product
    db:=ps.DB
     err:=db.Table("product").Find(&products)
    if err!=nil{
        log.Println(err)
    }
    return products,err
}

//查找商品剩余数量
func (ps *productService)GetProductNumById(ID int64)(int64,error){
    product:=model.Product{}
    db:=ps.DB
    _,err:=db.Table("product").Where("Id=?",ID).Get(&product)
    if err!=nil{
        log.Println(err)
    }
    return product.Productnum,err
}

//更新商品信息
func (ps *productService)UpdateServie(product model.Product)(err error){
    db:=ps.DB
    _,err=db.Table("product").Where("Id=?",product.Id).Update(&product)
    return
}

//商品数量减一
func (ps *productService)SubProductNum(productID int64)(err error){
    product:=model.Product{}
    db:=ps.DB
    _,err=db.Table("product").Where("Id=?",productID).Get(&product)
    if product.Productnum>=1{
        product.Productnum=product.Productnum-1
    }else{
        err=errors.New("商品售罄")
        return
    }
    _,err=db.Table("product").Where("Id=?",product.Id).Update(&product)
    return
}