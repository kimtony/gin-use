package controller

import (
	"encoding/json"
	"fmt"
	"gin-use/configs"
	"gin-use/src/global"
	"gin-use/src/util/consul"
	"gin-use/src/util/snowflake"
	"net/http"
	"time"

	"github.com/ddliu/go-httpclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var resp Resp


// Health 健康检查
// @Summary 健康检查接口
// @Description 服务是否启动正常检查
// @Tags 监测服务
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 200
// @Router /api/health [get]
func Health(c *gin.Context) {
	ip:=configs.GetLocalIp()
	now := time.Now()
	id := snowflake.GenerateId()
	workerid, datacenterid := snowflake.GetDeviceID(id)
	data := map[string]interface{}{"ip":ip,"Time":now,"id":id,"datacenterid": datacenterid,"workerid":workerid}
	Response("ok","请求成功",data,c)
}

func test()  {
	// 从consul中发现服务
	xichengCommon := consul.FindServer("xicheng-common","")
	global.Logger.Info("嘿，我能调用了 xicheng-common服务",zap.String("xichengCommon", xichengCommon))

	api := xichengCommon+"/api/health"

	res, err := httpclient.Get(api)
	if err != nil {
		fmt.Println("--------err----",err)
	}
	bodyString, err := res.ToString()
	if err != nil {
		fmt.Println("--------err----",err)
	}
	fmt.Println("--------bodyString----",bodyString)

	json.Unmarshal([]byte(string(bodyString)), &resp)
	fmt.Println("------------",resp.Code)
}


func Response(code string,msg string,data interface{},c *gin.Context) {
	c.JSON(http.StatusOK, Resp{
		code,
		msg,
		data,
	})
}
