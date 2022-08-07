package main

import (
	"Register-Login-Project/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run()

}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       //设置读取的文件名
	viper.SetConfigType("yml")               //设置读取文件类型
	viper.AddConfigPath(workDir + "/config") //设置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
