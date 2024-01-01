package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stable_diffusion_goweb/handle/aimage"
	"stable_diffusion_goweb/handle/album"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	// 设置CORS头，用于跨域请求
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080") // 允许特定域名访问
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "600")
		// axios
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	})
	//静态资源映射
	router.Static("/myAlbum", "./myAlbum")
	albumRouter := router.Group("/album")
	{
		//搜索图片
		albumRouter.GET("image", album.SearchImage)
		//修改图片
		albumRouter.POST("image", album.ModifyImage)
		//获取所有图片
		albumRouter.GET("/images", album.GetAllImages)
		// 前端保存图片到后端
		albumRouter.POST("/images", album.SaveImage)
		//批量删除图片
		albumRouter.POST("/delete", album.DeleteImages)
	}
	aimageRouter := router.Group("/aimage")
	{
		// 通过文本生成图片
		aimageRouter.POST("/txt2img", aimage.GetAiImage)
		//通过图片生成图片，接口其实是一样的
		aimageRouter.POST("/img2img", aimage.GetAiImage)
		//获取图片生成的进度
		aimageRouter.GET("/progress", aimage.GetAiProgress)
	}
	return router
}
