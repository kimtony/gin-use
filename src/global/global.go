package global

import (
	"gin-use/src/util/cache"
	"gin-use/src/util/db"

	"go.uber.org/zap"

	// "gin-use/configs"

	"github.com/spf13/viper"
)

var (
	DB     db.Repo
	Cache  cache.Repo
	Viper  *viper.Viper
	Logger *zap.Logger
)
