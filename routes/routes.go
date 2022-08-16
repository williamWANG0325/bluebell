package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"fmt"

	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		// 设置为发布模式，不输出调试信息
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return nil
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// swag
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 注册
	v1.POST("/signup", controller.SignupHandler)

	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatPostHandler)
		v1.GET("/post/:id", controller.GetPostDetail)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	return r
}
