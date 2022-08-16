package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"fmt"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed. err:%v\n", err)
		return
	}
	// 雪花算法，生成ID
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed. err:%v\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed. err:%v\n", err)
		return
	}
	// 3.初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed. err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4.初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed. err:%v\n", err)
		return
	}
	// 5.注册路由
	r := routes.Setup(settings.Conf.Mode)
	// 6.启动服务
	r.Run()
}
