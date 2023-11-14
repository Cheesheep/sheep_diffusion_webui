package model

import (
	"log"
	"stable_diffusion_goweb/initDB"
)

type Album struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	SavePath          string  `json:"save_path"`
	Height            int     `json:"height"`
	Width             int     `json:"width"`
	DenoisingStrength float32 `json:"denoising_strength"`
	ResizeMode        int     `json:"resize_mode"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	Steps             int     `json:"steps"`
	BatchSize         int     `json:"batch_size"`
	CfgScale          int     `json:"cfg_scale"`
	Seed              int     `json:"seed"`
	NIter             int     `json:"n_iter"`
	//OverrideSettings的两个属性直接取出来
	SDModelCheckpoint string `json:"sd_model_checkpoint"`
	SDVae             string `json:"sd_vae"`
	SamplerIndex      string `json:"sampler_index"`
}

// GetAll 获取所有图片
func (album Album) GetAll() []Album {
	var images []Album
	initDB.Db.Find(&images)
	return images
}
func (album Album) FindById() Album {
	initDB.Db.First(&album, album.Id)
	return album
}

func (album Album) Insert() int {
	create := initDB.Db.Create(&album)
	log.Println("图片添加成功，ID是", album.Id)
	if create.Error != nil {
		log.Panicln("图片添加到数据库失败！")
		return -1
	}
	return album.Id
}
func (album Album) Modify() {
	initDB.Db.Model(&Album{}).
		Where("id = ?", album.Id).
		Updates(Album{
			Name:   album.Name,
			Width:  album.Width,
			Height: album.Height,
		})
}

// Search 模糊搜索名字
func (album Album) Search(text string) []Album {
	var images []Album
	text = "%" + text + "%" //添加了%才能实现模糊匹配
	if err := initDB.Db.Where("name LIKE ?", text).Find(&images).Error; err != nil {
		log.Panicln("查询出错：", err)
	}
	return images
}
func (album Album) Delete() {
	initDB.Db.Delete(album)
}
