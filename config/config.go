package config

import(
    "os"
    "encoding/json"
)

type AppConfig struct{
    AppName  string   `json: "app_name"`
    Port     string   `json: "port"`
    DataBase DataBase `json:"data_base"`
    Rabbitmq Rabbitmq `json:"rabbit_mq"`
}

type DataBase struct{
    Driver   string `json:"driver"`
    User     string `json:"user"`
    Pwd      string `json:"pwd"`
    Host     string `json:"host"`
    Database string `json:"database"`
}

type Rabbitmq struct{
    User   string `json:"user"`
    Pwd    string `json:"pwd"`
    Url    string `json:"url"`
    Vhost  string `json:"vhost"`
}

//读取配置文件
func InitConfig()*AppConfig{
    file,err:=os.Open("./config/config.json")
    if err!=nil{
        panic(err.Error())
    }
    decoder:=json.NewDecoder(file)
    conf:=AppConfig{}
    err=decoder.Decode(&conf)
    if err!=nil{
        panic(err.Error())
    }
    return &conf
}