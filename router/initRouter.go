package router

import (
	"github.com/gin-gonic/gin"
	"stable_diffusion_goweb/handle/aimage"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	txt2imgRouter := router.Group("")
	{
		// 通过文本生成图片
		txt2imgRouter.POST("/txt2img", aimage.GetTxt2Image)
	}
	return router
}
