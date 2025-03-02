package routes

import (
	"GoWebX/logger"
	"GoWebX/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  routes.go.go
 * @Email: 1794748404@qq.com
 * @Date: 2025-02-28 17:56
 */

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(context *gin.Context) {
		context.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
