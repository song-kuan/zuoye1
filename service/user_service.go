package service

import(
    "seckill/model" 
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
	"log"
)

type UserService interface{
    IsLogIn(username,password string)(model.User,error)
    AddUser(user model.User)(userId int64,err error)
    IsUnique(username string)(bool)
}

type userService struct{
    DB *xorm.Engine
}

func NewUserService(db *xorm.Engine)UserService{
    return &userService{DB:db}
}

//判断用户是否登陆
func (us *userService)IsLogIn(username,password string)(model.User,error){
    user:=model.User{}
    db:=us.DB
    _,err:=db.Table("user").Where("username = ? and password = ?",username,password).Get(&user)
    if err!=nil{
        log.Println(err)
    }
    return user,err
}

//用户注册
func (us *userService)AddUser(user model.User)(int64,error){
    db:=us.DB
    _,err:=db.Table("user").Insert(&user)
    if err!=nil{
        log.Println(err)
    }
    return user.Id,err
}

//判断指定用户名是否存在，用于注册
func (us *userService)IsUnique(username string)(bool){
    user:=model.User{}
    db:=us.DB
    db.Table("user").Where("username = ?",username).Get(&user)
    if user.Id==0{
        return false
    } 
    return true
}