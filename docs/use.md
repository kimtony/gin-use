## 初始化使用
* go mod init 
* go mod tidy  |  go install
* swag init  
* sh scripts/start.sh


## go.mod
* 使用viper/remote包会报错需添加:
* replace google.golang.org/grpc => google.golang.org/grpc v1.28.0


## docker命令
* docker build -t docker.dev.isecsp.com/xicheng/car-test .
* docker push docker.dev.isecsp.com/xicheng/car-test


## viper配置
```
可读取多个配置文件
x := viper.New()
y := viper.New()
 
x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")
```

