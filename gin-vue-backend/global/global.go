package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GvaDb     *gorm.DB
	GvaDbList map[string]*gorm.DB
	GvaRedis  *redis.Client
	GvaVp     *viper.Viper
)
