package model

type TxtModel struct {
	Name             string `json:"name"`
	Prompt           string `json:"prompt"`
	NegativePrompt   string `json:"negative_prompt"`
	Steps            int    `json:"steps"`
	Height           int    `json:"height"`
	Width            int    `json:"width"`
	BatchSize        int    `json:"batch_size"`
	CfgScale         int    `json:"cfg_scale"`
	Seed             int    `json:"seed"`
	NIter            int    `json:"n_iter"`
	OverrideSettings struct {
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
