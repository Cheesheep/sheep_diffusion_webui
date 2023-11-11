# Stable Diffusion接口
### 文生图接口：

输入格式

~~~json
/sdapi/v1/txt2img
{
 "denoising_strength": 0,
 "prompt": "puppy dogs", //提示词
 "negative_prompt": "", //反向提示词
 "seed": -1, //种子，随机数
 "batch_size": 2, //每次张数
 "n_iter": 1, //生成批次
 "steps": 50, //生成步数
 "cfg_scale": 7, //关键词相关性
 "width": 512, //宽度
 "height": 512, //高度
 "restore_faces": false, //脸部修复
 "tiling": false, //可平埔
 "override_settings": {
     "sd_model_checkpoint" :"wlop-any.ckpt [7331f3bc87]"
}, // 一般用于修改本次的生成图片的stable diffusion 模型，用法需保持一致
   "script_args": [
      0,
      true,
      true,
      "LoRA",
      "dingzhenlora_v1(fa7c1732cc95)",
      1,
      1
  ], // 一般用于lora模型或其他插件参数，如示例，我放入了一个lora模型， 1，1为两个权重值，一般只用到前面的权重值1
 "sampler_index": "Euler" //采样方法
}
~~~
输出格式

~~~json
{
    "images": [...], // 这里是一个base64格式的字符串数组，根据你请求的图片数量而定
    "parameters": {
       //此处为你输入的body
    },
   // 返回的图片的信息
    "info": "{\"prompt\": \"puppy dogs\", \"all_prompts\": [\"puppy dogs\", \"puppy dogs\"], \"negative_prompt\": \"\", \"all_negative_prompts\": [\"\", \"\"], \"seed\": 2404186668, \"all_seeds\": [2404186668, 2404186669], \"subseed\": 3290733804, \"all_subseeds\": [3290733804, 3290733805], \"subseed_strength\": 0, \"width\": 512, \"height\": 512, \"sampler_name\": \"Euler\", \"cfg_scale\": 7.0, \"steps\": 50, \"batch_size\": 2, \"restore_faces\": false, \"face_restoration_model\": null, \"sd_model_hash\": \"7331f3bc87\", \"seed_resize_from_w\": -1, \"seed_resize_from_h\": -1, \"denoising_strength\": 0.0, \"extra_generation_params\": {}, \"index_of_first_image\": 0, \"infotexts\": [\"puppy dogs\\nSteps: 50, Sampler: Euler, CFG scale: 7.0, Seed: 2404186668, Size: 512x512, Model hash: 7331f3bc87, Seed resize from: -1x-1, Denoising strength: 0.0, ENSD: 31337\", \"puppy dogs\\nSteps: 50, Sampler: Euler, CFG scale: 7.0, Seed: 2404186669, Size: 512x512, Model hash: 7331f3bc87, Seed resize from: -1x-1, Denoising strength: 0.0, ENSD: 31337\"], \"styles\": [], \"job_timestamp\": \"20230422213724\", \"clip_skip\": 1, \"is_using_inpainting_conditioning\": false}"
}
~~~



### 后期处理接口：

![img](https://gitee.com/cheesheep/typora-photo-bed/raw/master/Timg/772544-20230831090155480-752232399.png)

输入格式

```json
{
  "resize_mode": 0,
  "show_extras_results": true,
  "gfpgan_visibility": 0,
  "codeformer_visibility": 0,
  "codeformer_weight": 0,
  "upscaling_resize": 2,
  "upscaling_resize_w": 512,
  "upscaling_resize_h": 512,
  "upscaling_crop": true,
  "upscaler_1": "None",
  "upscaler_2": "None",
  "extras_upscaler_2_visibility": 0,
  "upscale_first": false,
  "image": ""                       // 原图
}
```





## 后端问题解决

> 记录一些调试过程中遇到过的问题

### 1. 跨域请求

这个花了我好一段时间，是第一次链接前后端请求的时候

- 一开始遇到的问题是报错里面有`Error Network`之类的，这个实际上是由于我端口没有设置好

  例如前端这里是在main.js里面设置就可以了

  ~~~javascript
  axios.defaults.baseURL='http://172.29.35.238:9000'  
  ~~~

  然后后端需要设置CORS头

  ```go
  	router := gin.Default()
  	// 设置CORS头，用于跨域请求
  	router.Use(func(c *gin.Context) {
  		//c.Header("Access-Control-Allow-Origin", "http://localhost:8080")     // 允许特定域名访问
  		c.Header("Access-Control-Allow-Origin", "http://localhost:8080") // 允许特定域名访问
  		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
          .......
      }
  ```

  而我打开页面使用的是实验室里面的电脑，因此和我跑前后端的放在宿舍的电脑端口自然就不一样了，后端写成localhost，而前端这里又是网络端口，因此发生了错误

  如果是实验室电脑访问，那就应该前后端的IP都写成172开头的那个

- 解决了这个问题后，我发送空的POST请求成功了，但是发送带json数据的POST请求却**失败**了！

  这是为什么呢，查看报错，发现是跨域链接的过程中报错`Access to XMLHttpRequest at 'http://localhost:9000/txt2img' from origin 'http://localhost:8080' has been blocked by CORS policy: Response to preflight request doesn't pass access control check: It does not have HTTP ok status`

  前面一大堆的不用看，只看最后，It does not have HTTP ok status，然后我们再查看后端的控制台信息，是404的报错！

  ![image-20231025223852826](https://gitee.com/cheesheep/typora-photo-bed/raw/master/Timg/image-20231025223852826.png)

  也就是说其实数据已经发过去了，但是那边没有返回正确的内容，而且呢这个提交的方法是OPTION，这不是POST呀，我提交的明明是POST方法呀

  果然上网一搜，原来前端在发送POST请求前会先发一个OPTION请求，用来验证一下是否可以正常和后端通信，若不能POST请求就不发送了，这样也可以节约资源，嗯嗯，跟http的握手有点像呢哈哈

  于是添加相应的代码，注意是在`router.Use`当中喔

  ```go
  	router.Use(func(c *gin.Context) {
  		c.Header("Access-Control-Allow-Origin", "http://172.29.35.238:8080") // 允许特定域名访问
  		........
  		// axios
  		if c.Request.Method == http.MethodOptions {
  			c.JSON(http.StatusOK, nil)
  		}
  		c.Next()
  	})
  ```

  然后就可以愉快地进行前后端通信啦！

  ![image-20231025224318360](https://gitee.com/cheesheep/typora-photo-bed/raw/master/Timg/image-20231025224318360.png)

  





## 前端问题解决：

> 记录前端遇到过的一些值得记录的问题
