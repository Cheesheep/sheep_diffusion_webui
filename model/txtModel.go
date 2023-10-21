package model

type TxtModel struct {
	Prompt           string `form:"prompt"`
	NegativePrompt   string `form:"negative_prompt"`
	Step             int    `form:"step"`
	Height           int    `form:"height"`
	Width            int    `form:"width"`
	BatchSize        int    `form:"batch_size"`
	CfgScale         int    `form:"cfg_scale"`
	OverrideSettings struct {
		SDModelCheckpoint string `form:"sd_model_checkpoint"`
	} `form:"override_settings"`
	SamplerIndex string `form:"sampler_index"`
}
