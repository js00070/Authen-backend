package server

import (
	"authen/api"
	"authen/middleware"

	"github.com/gin-gonic/gin"
)

// Router 路由配置
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.TokenAuth())
	r.POST("/user/login", api.UserLogin)
	r.POST("/user/register", api.UserRegister)
	return r
}
