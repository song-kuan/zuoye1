package datasource

import(
    _ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"seckill/model"
	"seckill/config"
	"log"
)

func NewEngine()*xorm.Engine{
    //根据config加载数据库
    Config:=config.InitConfig()
    database:=Config.DataBase
    //user:pwd@/dbname?charset=utf8&&parseTime=true
    driverName:=database.Driver
    dataSource:=database.User+":"+database.Pwd+"@tcp("+database.Host+")/"+database.Database+"?charset=utf8"
    db,err:=xorm.NewEngine(driverName,dataSource)
    if err!=nil{
        log.Fatal(NewEngine, err)
	    return nil
    }
    //根据model同步数据库
    db.Sync2(new(model.Order),new(model.User),new(model.Product))
    //db.ShowSQL(true)
    return db
}