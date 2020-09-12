package Helpers

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"os"
)

type CodeStruct struct {
	Code string `json:"code"`
}

func SendSMS(mobile string ,code string)  string{
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ALIYUN_ACCESS_KEY_ID"), os.Getenv("ALIYUN_ACCESS_KEY_SECRET"))

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	codeStruct := CodeStruct{Code:code}
	codeJson,_ := json.Marshal(codeStruct)

	request.PhoneNumbers = mobile
	request.SignName = "小吉谱社区"
	request.TemplateCode = "SMS_185840965"
	request.TemplateParam = string(codeJson)
	//request.TemplateParam = "{\"code\":\"1111\"}"

	response, err := client.SendSms(request)
	if err != nil {
		LogInfo("sms.log" ,"短信获取response失败:"+err.Error())
		return "短信获取response失败"
	}

	if response == nil {
		LogInfo("sms.log" ,"短信获取response是nil")
		return "短信获取response是nil"
	}

	resMessage := response.Message
	resCode := response.Code

	if resMessage=="OK" && resCode=="OK" {
		return ""
	}else {
		LogInfo("sms.log" ,"短信发送失败，触发警告，:"+resMessage)
		return resMessage
	}
}


