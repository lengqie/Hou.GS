package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数

	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// ascii Json
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	// html 渲染
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// api参数
	r.GET("/api/:k1/*k2", func(c *gin.Context) {
		k1 := c.Param("k1")
		k2 := c.Param("k2")
		k2 = strings.Trim(k2, "/")
		//print(k1,k2)
		data := map[string]string{"k1":k1,"k2":k2}
		c.JSON(http.StatusOK, data)
	})
	// url 参数
	r.GET("/url", func(c *gin.Context) {
		query := c.Query("k")
		query = c.DefaultQuery("K","defaultValue")
		print(query)
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}