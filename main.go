package main

import (
	"nindychat/config"
	"nindychat/external"
	"nindychat/utils"
)

func main() {
	utils.LoadEnv(".env")

	external.InitializeGeminiEnv()
	external.InitializeGemini()

	config.InitializeBot()
}
