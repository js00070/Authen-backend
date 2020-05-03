package server

import (
	"authen/api"
	"authen/middleware"

	"github.com/gin-gonic/gin"
)

// Router 路由配置
func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/user/register", api.UserRegister)
	h := api.FileUpload()
	r.Any("/file/upload/*a", h)

	v := r.Group("")
	v.Use(middleware.TokenAuth())
	v.POST("/user/login", api.UserLogin)
	v.POST("/file/hash", api.FileHashUpload)
	v.GET("/filelist", api.GetFileList)
	return r
}
