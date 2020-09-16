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
		// fmt.Println(path)
		// WebUI展示相关的就跳过
		if path != "/" && !strings.Contains(path, "WebUI") {
			_key := ctx.Query("_key")
			if _key != Key {
				ctx.JSON(http.StatusBadRequest, Response{Ret: 1, Msg: "Faker!"})
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
