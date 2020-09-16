package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthorizeMiddleware 鉴权
// 配置里硬编码key
func AuthorizeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		host := strings.Split(ctx.Request.Host, ":")[0]
		if host == "127.0.0.1" || host == "0.0.0.0" || host == "localhost" || path == "/" || strings.Contains(path, "WebUI") {
			ctx.Next()
		} else {
			if _key := ctx.Query("_key"); _key == Key {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusBadRequest, Response{Ret: 1, Msg: "Faker!"})
				ctx.Abort()
			}
		}
	}
}
