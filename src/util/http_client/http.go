package util

import (
	"bytes"
	"encoding/json"
	"gin-use/src/global"
	"io/ioutil"
	"net/http"
)

// AllowOriginWriteJson ;跨域发送json数据
func AllowOriginWriteJson(w http.ResponseWriter, arg interface{}) (int, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := json.Marshal(arg)
	if err != nil {
		global.Logger.Errorf("json marshal rsp:%v fail, err:%v", arg, err)
		return 0, err
	}
	global.Logger.Debugf("write back:%s", string(data))
	return w.Write(data)
}


// 发送http post请求(json)
func HttpGet(addr string, arg interface{}) ([]byte, error) {
	data, err := json.Marshal(arg)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("GET", addr, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(rsp.Body)
}



// 发送http post请求(json)
func HttpPostWithJson(addr string, arg interface{}) ([]byte, error) {
	data, err := json.Marshal(arg)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", addr, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(rsp.Body)
}
