package validator

import (
	"encoding/json"
	// "reflect"
	"fmt"
	"strconv"
	"errors"
	"gin-use/src/controller"
	"gin-use/src/global"
	"strings"
    "github.com/gin-gonic/gin/binding"
    zhongwen "github.com/go-playground/locales/zh"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)
var v *validator.Validate
var trans ut.Translator


func InitVali() {
    // 中文翻译
    zh := zhongwen.New()
    uni := ut.New(zh, zh)
    trans, _ = uni.GetTranslator("zh")

    v, ok := binding.Validator.Engine().(*validator.Validate)
    if ok {
        // 验证器注册翻译器
        zh_translations.RegisterDefaultTranslations(v, trans)
        // 自定义验证方法
        v.RegisterValidation("checkMobile", checkMobile)
    }
}


func checkMobile(fl validator.FieldLevel) bool {
    mobile := strconv.Itoa(int(fl.Field().Uint()))
    if len(mobile) != 11 {
        return false
    }
    return true
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}



func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBind(request)
	var errStr string
	if err != nil {
		switch err.(type) {
		case *validator.ValidationErrors:
			errStr = Translate(err.(validator.ValidationErrors))
		case *json.UnmarshalTypeError :
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = errors.New(err.Error()).Error()
		}
		errStr = Translate(err.(validator.ValidationErrors))
		global.Logger.Error(err)
		controller.ResponseHttpOK("PARAMS_IS_INVALID", errStr,nil, c)
		return err
	}
	return nil
}


//Translate 翻译错误信息
func TranslateOtherErr(err error) map[string][]string {
	var result = make(map[string][]string)
 
	errors := err.(validator.ValidationErrors)
 
	for _, err := range errors{
	   result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
 }