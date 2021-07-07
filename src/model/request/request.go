package request

type Account struct {
	ID     	   string 	`json:"id" binding:"required"`
	Name   	   string 	`json:"name" binding:"required"`
	Mobile     string 	`json:"mobile"`
}


type WechatAccount struct {
	OpenID     string	 `json:"openid" binding:"required" `
	NickName   string  	 `json:"nickname" binding:"required"`
	Mobile 	   uint   	 `json:"mobile" binding:"required,gte=6"`
}

type GetNovelData struct {
	Cid 	  string     `json:"cid" form:"cid" binding:"required"`
	Num  	  uint	    `json:"num" form:"num" binding:"required"`
	Page 	  uint 		`json:"page" form:"page" binding:"required"`
}

