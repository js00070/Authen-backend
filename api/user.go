package api

import (
	"authen/service"

	"github.com/gin-gonic/gin"
	"github.com/ngaut/log"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterSvc
	err := c.BindJSON(&service)
	if err == nil {
		_, err := service.Register()
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "注册失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "注册成功",
				// "data": user,
			})
		}
	} else {
		log.Error(err)
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "json解析失败",
		})
	}

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginSvc

	if err := c.BindJSON(&service); err == nil {
		user, err := service.Login()
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "账号或密码不正确",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "登录成功",
				// "data":  user,
				"token": user.GetToken(),
			})
		}
	} else {
		log.Error(err)
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "json解析失败",
		})
	}
}
