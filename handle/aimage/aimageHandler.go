package aimage

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"stable_diffusion_goweb/model"
	"strings"
)

func SaveTxt2Image(context *gin.Context) {
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
	jsonData := `{
		"prompt": "NSFW, front angle, (8k, best quality, masterpiece:1.2), (realistic, photo-realistic:1.37), ultra-detailed, 1 girl, looking at viewer, beautiful detailed sky, detailed cafe street, sitting, full body, small head, intricate choker, (pretty legs:1.2), (long legs:1.2), slim legs, (high heels:1.3), (bare legs:1.4), medium breasts, high-waist, narrow waist, off-shoulder, belt, short bottoms, beautiful detailed eyes, daytime, warm tone, white lace, (long hair:1.4), silver medium hair, white skin, cinematic light, street light, <lora:koreandolllikeness_V10:0.2>, <lora:chineseGirl_v10:0.2>, <lora:cuteGirlMix4_v10:0.3>, <lora:chilloutmixss_xss10:0.4>",
		"negative_prompt": "easynegative, DreamArtistBADHAND, By bad artist -neg, (worst quality:2), (low quality:2), lowres, ((monochrome)), ((grayscale)), big head, severed legs:1.4, short legs, skin spots, acnes, skin blemishes, age spot, backlight,(ugly:1.331), (duplicate:1.331), (morbid:1.21), (mutilated:1.21), mutated hands, (poorly drawn hands:1.331), blurry, (bad anatomy:1.21), (bad proportions:1.331), (disfigured:1.331), (unclear eyes:1.331), bad hands, missing fingers, extra digit, bad body, NG_DeepNegative_V1_75T, pubic hair, glans",
		"step": 10,
		"height": 800,
		"width": 512,
		"batch_size": 2,
		"cfg_scale": 8,
		"override_settings": {
			"sd_model_checkpoint": "chilloutmix_NiPrunedFp32Fix.safetensors"
		},

		"sampler_index": "DPM++ SDE Karras"
	}`

	// 创建一个http.Client
	client := &http.Client{}

	// 创建一个POST请求
	req, err := http.NewRequest("POST", endpoint+"/sdapi/v1/txt2img", strings.NewReader(jsonData))
	if err != nil {
		log.Panicln("创建请求失败：", err)
		return
	}
	// 添加授权信息到请求头
	auth := username + ":" + password
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", authHeader)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("发送请求失败：", err)
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
		// 这里需要根据实际的JSON结构定义一个结构体来解析数据
		// 假设响应的JSON结构如下：
		type Response struct {
			Images []string `json:"images"`
			// 其他字段...
		}
		var data Response
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
				continue
			}
			err = os.WriteFile(fmt.Sprintf("album/%d.png", i), b, 0644)
			if err != nil {
				log.Panicln("保存图片失败：", err)
				continue
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
