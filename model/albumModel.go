package model

import (
	"log"
	"stable_diffusion_goweb/initDB"
)

type Album struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	SavePath string `json:"save_path"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	//Prompt         string `json:"prompt"`
	//NegativePrompt string `json:"negative_prompt"`
	//Steps          int    `json:"steps"`
	//BatchSize      int    `json:"batch_size"`
	//CfgScale       int    `json:"cfg_scale"`
	//Seed           int    `json:"seed"`
	//NIter          int    `json:"n_iter"`
	////OverrideSettings的两个属性直接取出来
	//SDModelCheckpoint string `json:"sd_model_checkpoint"`
	//SDVae             string `json:"sd_vae"`
	//SamplerIndex      string `json:"sampler_index"`
}

// GetAll 获取所有图片
func (album Album) GetAll() []Album {
	var images []Album
	initDB.Db.Find(&images)
	return images
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
			Name: album.Name,
		})
}
func (album Album) Delete() {
	initDB.Db.Delete(album)
}
