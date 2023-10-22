package model

type TxtModel struct {
	Prompt           string `json:"prompt"`
	NegativePrompt   string `json:"negative_prompt"`
	Step             int    `json:"step"`
	Height           int    `json:"height"`
	Width            int    `json:"width"`
	BatchSize        int    `json:"batch_size"`
	CfgScale         int    `json:"cfg_scale"`
	OverrideSettings struct {
		SDModelCheckpoint string `json:"sd_model_checkpoint"`
	} `json:"override_settings"`
	SamplerIndex string `json:"sampler_index"`
}

// 假设响应的JSON结构如下：
type TxtResponse struct {
	Images []string `json:"images"`
	// 其他字段...不是很重要..
	Info string `json:"info"`
}
