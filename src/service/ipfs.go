package service

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"gin-use/src/global"
	"io/ioutil"
	api "github.com/ipfs/go-ipfs-api"
)


//根据hash拿文件内容
func GetFileContent(addr string, hash string) ([]byte, error) {
	sh := api.NewShell(addr)
	read, err := sh.Cat(hash)
	if err != nil {
		global.Logger.Error("获取ipfs  报错:%v",err)
		return nil,err
	}
	body, err := ioutil.ReadAll(read)
	if err != nil {
		global.Logger.Error("获取ipfs 读取文件报错:%v",err)
		return nil,err
	}
	return body,nil
}

//将文件存入ipfs
func PutIpfsData(addr string, filename string) (string, error) {
	sh := api.NewShell(addr)
	data,err := ReadFile(filename)
	fmt.Println("-----filename-----",filename)
	if err != nil {
		fmt.Println("获取上传文件失败,错误:", err)
	}
	hash, err := sh.Add(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("放入文件失败,错误:", err)
		return hash, err
	}
	return hash, nil
}


func CurrPath() string {
	str, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return str
}

//读取本地文件内容
func ReadFile(name string) ([]byte, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}