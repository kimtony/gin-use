# gin-demo


# 后端技术栈
* 微服务开发
* 框架：node.js(eggs.js)+go(gin/mirco)
* 大前端业务中台：node 后端服务：go
* 服务注册与发现：consul
* API网关,流量转发统计：kong
* nginx:http+ssl


## hepler
* 生成uuid   Sonyflake
* 验证码
* 密码加密（hash）验证密码


## 数据库
* gorm+pgsql
* 简单增删改查询都用gorm
* 业务复杂用原生sql 
* 可以考虑postgrest接口统一由egg.js转发

## 缓存
* redis
* 缓存同步工具