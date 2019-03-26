package main

import (
	"fmt"
	"net/http"

	"Server/API5/config"
	"Server/API5/router"
	"Server/API5/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP(
		"config",             //名称
		"c",                  //速记
		"config/config.yaml", //配置文件路径
		"config file usage",
	)
)

func main() {
	/*
		在 main 函数中增加了 config.Init(*cfg) 调用，
		用来初始化配置，
		cfg 变量值从命令行 flag 传入，
		可以传值，
		比如 ./apiserver -c config.yaml，
		也可以为空，
		如果为空会默认读取 conf/config.yaml。
	*/
	pflag.Parse()

	err := config.Init(*cfg)
	if err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	engine := gin.Default()
	middlewares := []gin.HandlerFunc{}
	router.Load(engine, middlewares...)
	addr := viper.GetString("addr")
	fmt.Print(viper.GetString("alec.name"))
	http.ListenAndServe(addr, engine)
}


