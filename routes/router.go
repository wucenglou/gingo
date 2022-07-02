package routes

import (
	"net/http"

	"gingo/app/Http/Controllers/api"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("/some", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 简单的路由组: v1
	v1 := r.Group("/v1")
	{
		// v1.GET("/a", &api.BaseApi.Aa)
		v1.POST("/cc", api.Cc)
		v1.GET("/bb", api.Bb)
	}

	return r
}
