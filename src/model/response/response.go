package response


type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


type GetNovelDataPaginte struct {
	Sum   int  `json:"sum"`
	Data  string  `json:"data"`
}