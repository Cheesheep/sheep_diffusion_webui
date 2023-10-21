# Stable Diffusion接口
### 输入格式
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
### 输出格式

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

