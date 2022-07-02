package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Aa(c *gin.Context) {

	data := map[string]interface{}{
		"lang": "GO语言 aaaaaa",
		"tag":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

func Cc(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言 gogogo",
		"tag":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

func Bb(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言 gogogo dadada",
		"tag":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}
