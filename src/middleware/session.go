package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

/**
 * session中间件，
 * param: string keyPairs 要设置cookie的名字
 * return: gin.HandlerFunc
 */
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}

/**
 * 创建并返回一个sessions空间
 * return: sessions.Store
 */
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "xiaojipu"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}