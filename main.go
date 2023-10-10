package main

import (
	"ginchat/models"
	"ginchat/router" //  router "ginchat/router"
	"ginchat/utils"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	utils.InitConfig()
	// 初始化数据库
	utils.InitMySQL()
	// 初始化Redis
	utils.InitRedis()

	InitTimer()
	// 路由
	r := router.Router()                  // router.Router()
	r.Run(viper.GetString("port.server")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
