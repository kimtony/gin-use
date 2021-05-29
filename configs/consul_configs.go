package configs

import (
// "bytes"
// "fmt"
// "log"
// "time"

// consulapi "github.com/hashicorp/consul/api"
// "github.com/hashicorp/consul/api/watch"
// "github.com/spf13/viper"
// _ "github.com/spf13/viper/remote"
)

// var (
//     defaultConfig *viper.Viper
//     consulAddress string
//     consulPath    string
// )

// func initConfig() *viper.Viper {
//     consulAddress = "http://192.168.1.7/:8500"
//     consulPath = "test/test_config"

//     defaultConfig = viper.New()
//     defaultConfig.SetConfigType("toml")

//     consulClient, err := consulapi.NewClient(&consulapi.Config{Address: consulAddress})
//     if err != nil {
//         log.Fatalln("consul连接失败:", err)
//     }

//     kv, _, err := consulClient.KV().Get(consulPath, nil)
//     if err != nil {
//         log.Fatalln("consul获取配置失败:", err)
//     }

//     err = defaultConfig.ReadConfig(bytes.NewBuffer(kv.Value))
//     if err != nil {
//         log.Fatalln("Viper解析配置失败:", err)
//     }

//     go watchConfig()

//     return defaultConfig
// }
// func watchConfig() {
//     time.Sleep(time.Second * 10)
//     params := make(map[string]interface{})
//     params["type"] = "key"
//     params["key"] = consulPath

//     w, err := watch.Parse(params)
//     if err != nil {
//         log.Fatalln(err)
//     }
//     w.Handler = func(u uint64, i interface{}) {
//         kv := i.(*consulapi.KVPair)
//         hotconfig := viper.New()
//         hotconfig.SetConfigType("yaml")
//         err = hotconfig.ReadConfig(bytes.NewBuffer(kv.Value))
//         if err != nil {
//             log.Fatalln("Viper解析配置失败:", err)
//         }
//         defaultConfig = hotconfig
//     }
//     err = w.Run(consulAddress)
//     if err != nil {
//         log.Fatalln("监听consul错误:", err)
//     }
// }
// func GetConfig() *viper.Viper {
//     if defaultConfig == nil {
//         defaultConfig = initConfig()
//     }
//     return defaultConfig
// }
// func main() {
//     ReadOne()

//     go func() {
//         for {
//             host := GetConfig().GetString("store.bicycle.color")
//             fmt.Println("consul===", host)
//             time.Sleep(time.Second * 10)
//         }

//     }()

//     select {}
// }

// func ReadOne() {
//     runtimeConfig := viper.New()
//     runtimeConfig.AddRemoteProvider("consul", "http://192.168.100.19:8500", "config/v1/local")
//     runtimeConfig.SetConfigType("yaml")
//     err := runtimeConfig.ReadRemoteConfig()
//     if err != nil {
//         log.Fatalln("viper read:", err)
//     }
//     err = runtimeConfig.WatchRemoteConfigOnChannel()
//     if err != nil {
//         log.Fatalln("viper watch err:", err)
//     }
//     go func() {
//         for {
//             host := runtimeConfig.GetString("store.bicycle.color")
//             fmt.Println("viper=====", host)
//             time.Sleep(time.Second * 10)
//         }

//     }()

// }
