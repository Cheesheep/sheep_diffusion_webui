package main

import (
	"stable_diffusion_goweb/router"
)

func main() {
	router := router.SetUpRouter()
	_ = router.Run(":9000")
}
