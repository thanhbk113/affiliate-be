package initialize

import (
	"affiliate/internal/config"
	"affiliate/internal/locale"
)

func Init() {
	locale.LoadProperties()
	config.Init()
	RunCronJob()
	database()
	// Redis connect ...
	// redis.InitRedis()
}
