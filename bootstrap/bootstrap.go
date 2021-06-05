package bootstrap

import (

	"fmt"
	"gin-use/configs"
	"gin-use/src/util/cache"
	"gin-use/src/util/logger"
	"go.uber.org/zap"
	"gin-use/src/util/db"
	"gin-use/src/global"
	"gin-use/src/util/env"
	"gin-use/src/util/consul"

)


/**
系统初始化
*/
func Init() {


	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s:%s", configs.ProjectName(), env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	global.Logger = loggers
	defer loggers.Sync()

	// 初始化数据库
	dbRepo, err := db.New()
	if err != nil {
		loggers.Error("new db err", zap.Error(err))
	}
	global.DB = dbRepo


	//初始化缓存服务
	cacheRepo, err := cache.New()
	if err != nil {
		loggers.Error("new cahe err", zap.Error(err))
	}
	global.Cache = cacheRepo


	//consul服务注册与发现
	consul.Register()
	consul.CheckHeath()

	// //初始化定时任务
	// if config.GlobalConfig.CronTaskSwitch {
	// 	log.Logger.Info("系统初始化, 任务开关已开启")

	// 	//初始化全局定时任务调度器
	// 	cron.Init()

	// 	//加载定时任务列表
	// 	if err := s_cron.CronTaskService.LoadCronTaskList(); err != nil {
	// 		log.Logger.Error("系统初始化, 加载定时任务列表, 异常", zap.Error(err))
	// 	}
	// }
}
