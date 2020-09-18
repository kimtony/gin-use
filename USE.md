# 在使用的过程对比node和go的一些写法

## 场景: 构造一个对象参数
```
##node.js   直接写出来就行

const inpayload = {
    accountId : "123465",
    name : "lk",
    age : 18
}

##go写法 目前使用map(集合),可能还有别的写法   interface{}的作用就是你的对象值可以为自定义类型
##map[string]string  如果是这样定义则value都为string
    
inpayload := map[string]interface{}{
    "account": "3345472145214",
    "name":    "lk",
    "age":     18,
}

```

## 格式化输出 场景：构造rediskey时
```
fmt.Print：输出到控制台（仅只是输出）
fmt.Println：输出到控制台并换行
fmt.Printf：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）
fmt.Sprintf：格式化并返回一个字符串，不输出。
```
```
const id = 218910024463480831
redisKey := fmt.Sprintf("xicheng:account:%d", id)

#输出：xicheng:account:218910024463480833

```