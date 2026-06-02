package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/Dusheh/campus-market/internal/config"
	"github.com/Dusheh/campus-market/internal/handler"
	"github.com/Dusheh/campus-market/internal/middleware"
)

func Setup(db *gorm.DB, rdb *redis.Client, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())

	h := handler.NewHandler(db)

	api := r.Group("/api")
	{
		// 公开接口
		api.POST("/auth/login", h.Login)

		// 需要认证的接口
		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// 用户
			auth.GET("/user/profile", h.GetProfile)
			auth.PUT("/user/profile", h.GetProfile)

			// 服务
			auth.GET("/services", h.ListServices)
			auth.GET("/services/:id", h.GetService)
			auth.POST("/services", h.CreateService)

			// 物品
			auth.GET("/goods", h.ListGoods)
			auth.GET("/goods/:id", h.GetGoods)
			auth.POST("/goods", h.CreateGoods)

			// 需求
			auth.GET("/demands", h.ListDemands)
			auth.GET("/demands/:id", h.GetDemand)
			auth.POST("/demands", h.CreateDemand)

			// 订单
			auth.GET("/orders", h.ListOrders)
			auth.POST("/orders", h.CreateOrder)
			auth.PUT("/orders/:id/status", h.UpdateOrderStatus)
		}
	}

	return r
}