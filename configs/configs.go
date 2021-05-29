package configs

import (
	"fmt"
	"gin-use/src/util/env"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var build strings.Builder
var config = new(Config)

type Config struct {
	//数据库配置
	Pg struct {
		Read struct {
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Pass string `json:"pass"`
			Name string `json:"name"`
		} `json:"read"`
		Write struct {
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Pass string `json:"pass"`
			Name string `json:"name"`
		} `json:"write"`
		Base struct {
			MaxOpenConn     int           `json:"maxOpenConn"`
			MaxIdleConn     int           `json:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `json:"connMaxLifeTime"`
		} `json:"base"`
	} `json:"pg"`

	//redis缓存
	Redis struct {
		Addr         string `json:"addr"`
		Pass         string `json:"pass"`
		Db           int    `json:"db"`
		MaxRetries   int    `json:"maxRetries"`
		PoolSize     int    `json:"poolSize"`
		MinIdleConns int    `json:"minIdleConns"`
	} `json:"redis"`
	
	
	Mail struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		User string `json:"user"`
		Pass string `json:"pass"`
		To   string `json:"to"`
	} `json:"mail"`

	JWT struct {
		Secret         string        `json:"secret"`
		ExpireDuration time.Duration `json:"expireDuration"`
	} `json:"jwt"`

	URLToken struct {
		Secret         string        `json:"secret"`
		ExpireDuration time.Duration `json:"expireDuration"`
	} `json:"urlToken"`

	HashIds struct {
		Secret string `json:"secret"`
		Length int    `json:"length"`
	} `json:"hashids"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	// getConsul()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func getConsul() {
	var v = viper.New()

	v.AddRemoteProvider("consul", "http://192.168.1.7:8500", "test/test_config")

	v.SetConfigType("json") // 因为不知道格式，所以需要指定，支持的格式有"json"、"json"、"yaml"、"yml"、"properties"、"props"、"prop"

	if err := v.ReadRemoteConfig(); err != nil {
		fmt.Println("获取配置文件报错", err)
		return
	}

	fmt.Println("获取配置文件的port", v.GetInt("port"))
	fmt.Println("获取配置文件的mysql.url", v.GetString(`mysql.url`))
	fmt.Println("获取配置文件的mysql.username", v.GetString(`mysql.username`))
	fmt.Println("获取配置文件的mysql.password", v.GetString(`mysql.password`))
	fmt.Println("获取配置文件的redis", v.GetStringSlice("redis"))
	fmt.Println("获取配置文件的smtp", v.GetStringMap("smtp"))

	// 从远程读取配置
	// err := v.ReadRemoteConfig()

	// // 解析配置到runtime_conf中
	// runtime_viper.Unmarshal(&runtime_conf)

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
