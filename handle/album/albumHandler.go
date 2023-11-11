package album

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"stable_diffusion_goweb/handle/aimage"
	"stable_diffusion_goweb/model"
	"strconv"
	"strings"
	"time"
)

func SaveImage(context *gin.Context) {
	imageArray := context.PostFormArray("images")
	// 处理每张图片的Base64数据
	for i, imageBase64 := range imageArray {
		// 解码Base64数据
		imageBytes, err := base64.StdEncoding.DecodeString(imageBase64)
		if err != nil {
			context.JSON(400, gin.H{"message": "第" + strconv.Itoa(i) + "张图片的Base64解码失败"})
			return
		}
		// 创建唯一文件名称
		currentTime := time.Now()
		filePath := "myAlbum/image" + strconv.Itoa(i) +
			currentTime.Format("20060102150405") + ".png"
		//保存图片到本地
		err = os.WriteFile(fmt.Sprintf(filePath), imageBytes, 0644)
		if err != nil {
			log.Panicln("保存图片失败：", err)
		}
		//添加图片信息到数据库
		image := model.Album{}
		var id = -1 //默认id为-1
		image.SavePath = filePath
		//把路径的前缀和后缀去掉
		var fileName = strings.TrimSuffix(filePath, ".png")
		image.Name = strings.TrimSuffix(fileName, "myAlbum/")
		image.Width = aimage.GlobalTxt.Width
		image.Height = aimage.GlobalTxt.Height
		if id = image.Insert(); id == -1 {
			log.Panicln("保存图片失败：", err)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "保存图片成功",
	})
}

// GetAllImages 相册界面获取所有图片
func GetAllImages(context *gin.Context) {
	image := model.Album{}
	images := image.GetAll()
	context.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

func ModifyImage(context *gin.Context) {
	newImage := model.Album{}
	if e := context.ShouldBind(&newImage); e != nil {
		log.Panicln("绑定修改图片的数据失败!", e)
	}
	log.Println(newImage)
	//更新album的字段内容
	newImage.Modify()
	context.JSON(http.StatusOK, gin.H{
		"msg": "修改成功!",
	})
}
func SearchImage(context *gin.Context) {
	searchText := context.Query("searchInput")
	log.Println("搜索文本", searchText)
	image := model.Album{}
	images := image.Search(searchText)
	context.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

// DeleteImages 删除多张图片
func DeleteImages(context *gin.Context) {
	idArray := context.PostFormArray("selectedImages")
	log.Println("删除的数组：", idArray)
	for _, eachId := range idArray {
		if id, err := strconv.Atoi(eachId); err != nil {
			log.Panicln("ID转换失败", err)
		} else {
			image := model.Album{
				Id: id,
			}
			err = os.Remove(image.FindById().SavePath)
			if err != nil {
				log.Panicln("删除本地图片失败:", err)
			}
			image.Delete() //数据库删除相应的图片记录
			log.Println("第", id, "张图片删除成功！")
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "删除图片成功！",
	})
}
