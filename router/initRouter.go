package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stable_diffusion_goweb/handle/aimage"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	// 设置CORS头，用于跨域请求
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080") // 允许特定域名访问
		//c.Header("Access-Control-Allow-Origin", "http://172.29.35.238:8080") // 允许特定域名访问
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "600")
		// axios
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	})
	txt2imgRouter := router.Group("")
	{
		// 通过文本生成图片
		txt2imgRouter.POST("/txt2img", aimage.GetTxt2Image)
	}
	return router
}
