package model

type TxtModel struct {
	Name              string   `json:"name"`
	InitImages        []string `json:"init_images"`        //图生图的图片接口
	DenoisingStrength float32  `json:"denoising_strength"` //图生图重绘幅度
	ResizeMode        int      `json:"resize_mode"`
	Prompt            string   `json:"prompt"`          //正向提示词
	NegativePrompt    string   `json:"negative_prompt"` //反向提示词
	Steps             int      `json:"steps"`
	Height            int      `json:"height"`
	Width             int      `json:"width"`
	BatchSize         int      `json:"batch_size"`
	CfgScale          int      `json:"cfg_scale"`
	Seed              int      `json:"seed"`
	NIter             int      `json:"n_iter"`
	OverrideSettings  struct {
		SDModelCheckpoint string `json:"sd_model_checkpoint"`
		SDVae             string `json:"sd_vae"`
	} `json:"override_settings"`
	SamplerIndex string `json:"sampler_index"`
}

// TxtResponse 响应的JSON结构如下：
type TxtResponse struct {
	Images []string `json:"images"`
	// 其他字段...不是很重要..
	Info string `json:"info"`
}
