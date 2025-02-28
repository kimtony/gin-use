package configs

import (
	"fmt"
	"gin-use/src/util/env"
	"net"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

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

//初始化consul配置文件
func init() {
	var viper = viper.New()

	//config_path 路径为 consul下key value的文件路径
	viper.AddRemoteProvider("consul", ConsulAddr(), os.Getenv("CONSUL_CONFIG_PATH"))

	if os.Getenv("CONSUL_HOST") == "" {
		fmt.Println("获取consul配置CONSUL_HOST为空 无法启用consul服务注册与发现!")
	}

	//指定读取json文件
	viper.SetConfigType("json")

	//获取consul配置
	if err := viper.ReadRemoteConfig(); err != nil {
		fmt.Println("获取consul配置文件报错,启用本地配置", err)
		readLocalFileConfig()
		// panic(err)
	}

	//将json字符串解码到相应的数据结构
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	//开启一个goroutine来永远监听远程变化
	go func() {
		for {
			//delay after each request
			time.Sleep(time.Second * 10) //10秒监听一次
			err := viper.WatchRemoteConfig()
			if err != nil {
				fmt.Errorf("unable to read remote config: %v", err)
				continue
			}
			fmt.Println("从consul读取配置文件 时间:", time.Now())
			viper.OnConfigChange(func(e fsnotify.Event) {
				if err := viper.Unmarshal(config); err != nil {
					panic(err)
				}
			})
			//将当前配置 写入/覆盖 到自定义的路径中
			viper.SetConfigName(env.Active().Value() + "_configs")
			viper.SetConfigType("json")
			viper.AddConfigPath("./configs")
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("WriteConfig As Failed Error %s \n", err)
				panic(err)
			}
		}
	}()
}

//获取配置信息
func Get() Config {
	return *config
}

//机器id
func MachineId() string {
	return "raspi4b8g-demo-0001"
}

//获取项目名字
func ProjectName() string {
	return "gin-use"
}

//获取主机ip
func ProjectHost() string {
	return fmt.Sprintf("http://%s", GetLocalIp()[0])
}

//获取端口
func ProjectPort() string {
	return "8081"
}

func IpfsAddr() string {
	return "http://47.103.9.157:5001"
}

//获取consul地址
func ConsulAddr() string {
	return fmt.Sprintf("http://%s:%s", os.Getenv("CONSUL_HOST"), os.Getenv("CONSUL_PORT"))
}

//接口文档
func SwaggerUrl() string {
	return fmt.Sprintf("%s:%s%s", ProjectHost(), ProjectPort(), "/sys/swagger/doc.json")
}

//日志文件目录
func ProjectLogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}

//获取本地ip
func GetLocalIp() []string {
	var ipStrArr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces error:", err.Error())
		return ipStrArr
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv4
					if ipnet.IP.To4() != nil {
						ipStrArr = append(ipStrArr, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ipStrArr

}

//读取本地的配置文件
func readLocalFileConfig() {
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("json")
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
