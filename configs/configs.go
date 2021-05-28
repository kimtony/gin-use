package configs

import (
	"fmt"
	"strings"
	"time"

	"gin-use/src/util/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var build strings.Builder
var config = new(Config)

type Config struct {
	Pg struct {
		Read struct {
			Host string `toml:"host"`
			Port string `toml:"port"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Host string `toml:"host"`
			Port string `toml:"port"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"pg"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`

	Mail struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		To   string `toml:"to"`
	} `toml:"mail"`

	JWT struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"jwt"`

	URLToken struct {
		Secret         string        `toml:"secret"`
		ExpireDuration time.Duration `toml:"expireDuration"`
	} `toml:"urlToken"`

	HashIds struct {
		Secret string `toml:"secret"`
		Length int    `toml:"length"`
	} `toml:"hashids"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

//获取配置信息
func Get() Config {
	return *config
}

//获取项目名字
func ProjectName() string {
	return "gin-use"
}

//获取主机ip
func ProjectHost() string {
	return "http://192.168.1.163"
}

//获取端口
func ProjectPort() string {
	return ":8081"
}

//接口文档
func SwaggerUrl() string {
	
	build.WriteString(ProjectHost())
	build.WriteString(ProjectPort())
	build.WriteString("/sys/swagger/doc.json")
	return build.String()
}
func ProjectLogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}

func ProjectInstallFile() string {
	return "INSTALL.lock"
}
