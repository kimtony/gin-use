# gin-demo

```
## 初始化使用
* go mod init 
* go mod tidy  |  go install
* swag init  
* sh scripts/start.sh

## go.mod
* 使用viper/remote包会报错需添加:
* replace google.golang.org/grpc => google.golang.org/grpc v1.28.0

```


## 开发工具 vscode
* [插件+代码块](https://www.liwenzhou.com/posts/Go/00_go_in_vscode/)

## 接口文档
* [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 

## 后端技术栈
* 框架：go-gin (api + grpc)

## nginx | trafix 
* 反向代理与负载均衡
* api网关

## consul 
* 服务注册与发现
* 服务健康检查
* consul的key-value+viper 实现 动态配置值

## 日志
* zap日志模块
* loki+promatil+grafana

## 监控
* prometheus + grafana
* [Prometheus](https://github.com/prometheus/client_golang) 

## pprof
* [pprof](https://github.com/gin-contrib/pprof) 性能剖析 

## 数据库
* gorm+pgsql
* gorm文档：https://jasperxu.github.io/gorm-zh/advanced.html

## 缓存
* [go-redis](https://github.com/go-redis/redis/v7)


## util
* 生成uuid [雪花算法](https://juejin.cn/post/6844904035380658190)
* 验证码
* 密码加密（hash）验证密码
* jwt token

## dockerfile
* 多层缓存构建
* 减少镜像体积

## protobuf序列化
```
go get  -v -u github.com/golang/protobuf/proto
go get  -v -u github.com/golang/protobuf/protoc-gen-go
```

## devops
* docker部署
* gitlab ci cd


## git flow
* git cz模块
* git remote -v 查看远程仓库地址
* git branch 查看当前分支

## 微信相关
* [微信开发](https://silenceper.com/wechat/officialaccount/start.html)
* 微信小程序
* 微信公众号
* 微信支付

## 代码安全
* [go代码安全](https://github.com/Tencent/secguide/blob/main/Go%E5%AE%89%E5%85%A8%E6%8C%87%E5%8D%97.md)


## 其他
```
## viper配置

可读取多个配置文件
x := viper.New()
y := viper.New()
 
x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")


##构造一个对象参数

map[string]string  如果是这样定义则value都为string
inpayload := map[string]interface{}{
    "account": "3345472145214",
    "name":    "lk",
    "age":     18,
}
```