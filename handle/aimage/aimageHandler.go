package aimage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"stable_diffusion_goweb/model"
)

func GetTxt2Image(context *gin.Context) {
	log.Println("开始生成图片")
	endpoint := "http://sd.sheep-aliyun-stable-diffusion-webui.1770045088640528.ap-northeast-1.fc.devsapp.net"
	username := "sheep"
	password := "110120"

	var txt model.TxtModel
	if err := context.ShouldBind(&txt); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误: ", err.Error())
	}
	// 构建请求的JSON数据
	jsonData, err := json.Marshal(txt)
	if err != nil {
		log.Panicln("转换json格式失败", err.Error())
	}

	// 创建一个http.Client
	client := &http.Client{}

	// 创建一个POST请求
	req, err := http.NewRequest("POST", endpoint+"/sdapi/v1/txt2img", bytes.NewReader(jsonData))
	if err != nil {
		log.Panicln("创建请求失败：", err.Error())
		return
	}
	// 添加授权信息到请求头
	auth := username + ":" + password
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", authHeader)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("发送请求失败：", err.Error())
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panicln("读取响应数据失败：", err)
			return
		}

		// 解析响应JSON数据
		var data model.TxtResponse
		//将json格式的body转换成结构体到data当中
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Panicln("解析响应数据失败：", err)
			return
		}

		// 将图片解析出来(base64)并保存
		for i, img := range data.Images {
			b, err := base64.StdEncoding.DecodeString(img)
			if err != nil {
				log.Panicln("解码图片失败：", err)
			}
			err = os.WriteFile(fmt.Sprintf("album/%d.png", i), b, 0644)
			if err != nil {
				log.Panicln("保存图片失败：", err)
			}
		}
	} else {
		log.Panicln("请求失败，状态码：", resp.StatusCode)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":   "生成图像成功！",
		"image": "图片",
	})
}
