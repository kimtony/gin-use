package v1

import (
	"os"
	"io"
	"fmt"
	"strconv"
	// "gin-use/src/global".
	"gin-use/src/util/validator"
	"gin-use/src/controller"
	"gin-use/src/service"
	"gin-use/configs"
	"gin-use/src/model/request"
	"gin-use/src/model/response"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"

)


var (
	reqGetNovelData request.GetNovelData
)
// GetNovelData 获取小说数据
// @Summary 根据hash下载ipfs数据
// @Description 小说dapp小说文本数据
// @Tags 小说
// @Accept application/json
// @Produce application/json
// @Param cid query string false "string valid"
// @Param num query int false "int valid"
// @Param page query int false "int valid"
// @Success 200 {object} response.Resp
// @Router /v1/ipfs/novel/get [get]
func GetNovelData(c *gin.Context) {


	//参数校验
	if err := validator.ParseRequest(c,&reqGetNovelData); err != nil {
        return
    }
	cid 	:= c.Query("cid")

	num  ,_ := strconv.Atoi(c.Query("num"))
	page ,_ := strconv.Atoi(c.Query("page"))
	
	ipfsData,err := service.GetFileContent(configs.IpfsAddr(),cid)
	if err != nil {
	 	controller.ResponseHttpOK("fail", "请求失败!", "", c)
	}else{
		//数据处理
		data := string(ipfsData);
		sum  := len([]rune(data))/num
		subData := SubString(data,(page-1)*num,page*num)
		res := &response.GetNovelDataPaginte{
			Sum  : sum,
			Data : subData,
		}
		controller.ResponseHttpOK("ok", "请求成功", res , c)
	}
}



// PutNovelData 上传小说数据到ipfs
// @Summary 上传ipfs数据
// @Description 上传小说文本章节,文本类型txt等上传成功返回文本hash 并且需要记录一下数据到pg 小说标题,图片,作者,上传到ipfs的hash。
// @Tags 小说
// @Accept application/json
// @Produce application/json
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Success 200 {object} response.Resp
// @Router /v1/ipfs/novel/upload [post]
func PutNovelData(c *gin.Context) {
	//获取文件头
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		controller.ResponseHttpOK("fail", "请求失败!", "", c)
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename
	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	tmp, err := os.Create(filename)
	if err != nil {
		fmt.Println("报错:", err)
	}
	defer tmp.Close()

	//中文转码
	decoder := mahonia.NewDecoder("gbk")
	// 将file的内容拷贝到tmp
	_, err = io.Copy(tmp, decoder.NewReader(file))
	if err != nil {
		fmt.Println("报错:", err)
	}
	
	hash,err := service.PutIpfsData(configs.IpfsAddr(),filename)
	if err != nil {
	 	controller.ResponseHttpOK("fail", "请求失败!", "", c)
	}else{
		// 删除文件
		if err := os.Remove(filename); err != nil {
			fmt.Println("上传ipfs文件后删除本地文件 报错:",err)
		}
		controller.ResponseHttpOK("ok", "请求成功", hash, c)
	}
}




//用于截取一段字符串的函数msubstr()
func SubString(str string, begin int, length int) string {
	// fmt.Println("Substring =", str)
	rs := []rune(str)
	lth := len(rs)
	// fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	// fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	return string(rs[begin:end])
}







// //定义一个用于截取一段字符串的函数msubstr()
// func msubstr(str string,start int,len int) string {    	//str指的是字符串,$start指的是字符串的起始位置，$len指的是长度。
// 	strlen=start+len;								//用$strlen存储字符串的总长度（从字符串的起始位置到字符串的总长度）
// 	for i := 0; i <= strlen; i++ {					//通过for循环语句,循环读取字符串
// 		if(ord(substr(str,i,1))>0xa0){   			//如果字符串中首个字节的ASCII序数值大于0xa0,则表示为汉字
// 			tmpstr.=substr(str,i,2);				//每次取出两位字符赋给变量$tmpstr，即等于一个汉字
// 			i++;									//变量自加1
// 	   }else{										//如果不是汉字，则每次取出一位字符赋给变量$tmpstr
// 			 tmpstr.=substr(str,i,1);
// 	   }
// 	}
// 	return tmpstr;									//输出字符串
// }

