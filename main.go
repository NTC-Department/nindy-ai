package main

import (
	"nindychat/config"
	"nindychat/database/redis"
	"nindychat/external"
	"nindychat/utils"
)

func main() {
	utils.LoadEnv(".env")

	external.InitializeGemini()
	redis.InitializeRedis()
	config.InitializeBot()
}
