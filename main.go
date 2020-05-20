package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"github.com/ziqiango/goadmin/core"
	"github.com/ziqiango/goadmin/web/admin/routes"
	"os"
)

func main() {

	// 加载配置
	viper.SetConfigName(".config") // name of config file (without extension)
	viper.SetConfigType("yml") // REQUIRED if the config file does not have the extension in the name
	path, err := os.Getwd()
	viper.AddConfigPath(path)   // path to look for the config file in
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 实例化iris
	app := iris.New()

	app.Logger().SetLevel(viper.GetString("log.level"))


	// 加载路由
	routes.Configure(app)

	// 注册mysql

	core.MysqlDriveInit()


	// 注册redis


	// 启动iris

	app.Run(iris.Addr(viper.GetString("app.web")))
}
