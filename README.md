# gin-demo

# 运行说明
* go mod init 
* go install
* ./start.sh
* 如果找不到这个包go get github.com/silenceper/gowatch


# go相关资料文档
* https://www.bookstack.cn/read/topgoer/b0a74e6ce3f8548b.md

# 后端技术栈
* 微服务开发
* 框架：node.js(eggs.js)+go(gin/mirco)
* 大前端业务中台：node 后端服务：go
* 服务注册与发现：consul
* API网关,流量转发统计：kong
* nginx:http+ssl
* docker部署


## hepler
* 生成uuid   Sonyflake
* 验证码
* 密码加密（hash）验证密码
* jwt token

## middleware
* servicedefine校验

## 数据库
* gorm+pgsql
* 简单增删改查询都用gorm
* 业务复杂用原生sql 
* 可以考虑postgrest接口统一由egg.js转发
* gorm文档：https://jasperxu.github.io/gorm-zh/advanced.html


## 缓存
* redis
* 缓存同步工具


## 平滑启动
* gowatch :go get github.com/silenceper/gowatch


## sentry 
* http://192.168.1.10:9000/auth/login/sentry/
* yuezhi@qq.com   pass@2020





