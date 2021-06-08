package controller

import (
	"encoding/json"
	"fmt"
	"gin-use/configs"
	"gin-use/src/global"
	"gin-use/src/model/response"
	"gin-use/src/util/consul"
	"gin-use/src/util/snowflake"
	"net/http"
	"reflect"
	"time"

	"github.com/ddliu/go-httpclient"
	"github.com/gin-gonic/gin"
)

var resp *response.Resp

// Health 健康检查
// @Summary 健康检查接口
// @Description 服务是否启动正常检查
// @Tags 监测服务
// @Accept application/json
// @Produce application/json
// @Param name query string false "用户名"
// @Success 200
// @Router /api/health [get]
func Health(c *gin.Context) {
	ip := configs.GetLocalIp()
	now := time.Now()
	id := snowflake.GenerateId()
	workerid, datacenterid := snowflake.GetDeviceID(id)
	data := map[string]interface{}{"ip": ip, "Time": now, "id": id, "datacenterid": datacenterid, "workerid": workerid}
	Response("ok", "请求成功", data, c)
}

func test() {
	// 从consul中发现服务
	xichengCommon := consul.FindServer("xicheng-common", "")
	global.Logger.Info("嘿，我能调用了 xicheng-common服务, xichengCommon:%s", xichengCommon)

	api := xichengCommon + "/api/health"

	res, err := httpclient.Get(api)
	if err != nil {
		global.Logger.Errorf("http-client, err:%v", err)
	}
	bodyString, err := res.ToString()
	if err != nil {
		fmt.Println("--------err----", err)
	}
	fmt.Println("--------bodyString----", bodyString)

	json.Unmarshal([]byte(string(bodyString)), &resp)
	fmt.Println("------------", resp.Code)
}

//api响应值
func Response(code string, msg string, data interface{}, c *gin.Context) {
	//反射判断interface是否为空值
	if reflect.TypeOf(data) != nil {
		c.JSON(http.StatusOK, response.Resp{ code,msg,data })
	}

	c.JSON(http.StatusOK, response.Resp{ code,msg,map[string]interface{}{} })

}
