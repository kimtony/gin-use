> go相关使用
## [Golang 调度器 GMP 原理与调度全分析](https://learnku.com/articles/41728)
## go并发
* goroutine、channel、mutex、syncmap

## 场景: 构造一个对象参数
```
##map[string]string  如果是这样定义则value都为string
inpayload := map[string]interface{}{
    "account": "3345472145214",
    "name":    "lk",
    "age":     18,
}

```

## 字符串和各种int类型之间的相互转换方式：
```
string转成int：
int, err := strconv.Atoi(string)
string转成int64：
int64, err := strconv.ParseInt(string, 10, 64)
int转成string：
string := strconv.Itoa(int)
int64转成string：
string := strconv.FormatInt(int64,10)
```


## 字符串拼接
```
### 性能最好,官方建议
s1 := "字符串"
s2 := "拼接"
var build strings.Builder
build.WriteString(s1)
build.WriteString(s2)
s3 := build.String()
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