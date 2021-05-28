package global

import (
	"gin-use/src/util/cache"
	"gin-use/src/util/db"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	DB     db.Repo
	Cache  cache.Repo
	Viper  *viper.Viper
	Logger *zap.Logger
)
