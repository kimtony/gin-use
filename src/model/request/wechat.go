package request

type WechatAccount struct {
	OpenID     string `json:"openid" binding:"required"`
	NickName   string `json:"nickname" binding:"required"`
	Mobile 	   uint   `json:"mobile" binding:"required,gte=6"`
}
