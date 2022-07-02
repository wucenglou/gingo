package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Base struct{}

func (b *Base) Login(c *gin.Context) {

	data := map[string]interface{}{
		"lang": "GO语言 gogogo",
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
