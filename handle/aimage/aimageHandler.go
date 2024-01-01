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
	"net/url"
	"stable_diffusion_goweb/model"
)

var GlobalTxt model.TxtModel

func GetAiImage(context *gin.Context) {
	log.Println("开始生成图片")
	//endpoint := "http://cheesheep.xyz"
	endpoint := "http://sd.fc-stable-diffusion-plus.1770045088640528.cn-shenzhen.fc.devsapp.net"
	username := "sheep"
	password := "110120"
	// 解析响应JSON数据
	var data model.TxtResponse
	var txt model.TxtModel
	//接收json数据，转成结构体（不直接使用是为了保存传入的数据，并且方便观看）
	if err := context.ShouldBind(&txt); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误: ", err.Error())
	}
	GlobalTxt = txt //更改全局提示词信息，表示当前提示词已经发生改变
	log.Println(txt.Name, " 前端传输的文本信息：", txt.DenoisingStrength)
	log.Println(txt.Name, " 前端传输的文本信息：", txt.ResizeMode)
	log.Println(txt.Name, " 前端传输的文本信息：", txt.Seed)
	// 将结构体转成Json
	jsonData, err := json.Marshal(txt)
	if err != nil {
		log.Panicln("转换json格式失败", err.Error())
	}

	// 创建一个http.Client
	client := &http.Client{}
	req := new(http.Request)
	// 创建一个POST请求
	if txt.Name == "txt2img" {
		req, err = http.NewRequest("POST", endpoint+"/sdapi/v1/txt2img", bytes.NewReader(jsonData))
	} else if txt.Name == "img2img" {
		req, err = http.NewRequest("POST", endpoint+"/sdapi/v1/img2img", bytes.NewReader(jsonData))
	}
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
	//ReadReceiveData(resp, &data)
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panicln("读取响应数据失败：", err)
			return
		}

		//将json格式的body转换成结构体到data当中
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Panicln("解析响应数据失败：", err)
			return
		}
		//返回成功结果
		context.JSON(http.StatusOK, gin.H{
			"msg":    "生成图像成功！",
			"images": data.Images,
		})
	} else {
		//返回错误结果
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "云端解析图像数据失败！",
			"images": data.Images,
		})
		log.Panicln("云端解析图像数据失败，状态码：", resp.StatusCode)
	}
}

// GetAiProgress 获取图片生成进度
func GetAiProgress(context *gin.Context) {
	log.Println("获取图片生成进度")
	// 目标URL
	targetURL := "http://sd.fc-stable-diffusion-plus.1770045088640528.cn-shenzhen.fc.devsapp.net/sdapi/v1/progress"

	// 构建Query参数
	queryParams := url.Values{}
	queryParams.Set("skip_current_image", "false") // 默认为false

	// 构建完整的URL，包括Query参数
	fullURL := targetURL + "?" + queryParams.Encode()

	// 创建GET请求
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		return
	}

	// 添加Authorization头部
	username := "sheep"
	password := "110120"
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	request.Header.Add("Authorization", authHeader)

	// 发送GET请求
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 解析JSON数据
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 打印JSON数据
	fmt.Println("Response Data:")
	fmt.Printf("%+v\n", responseData)
}
