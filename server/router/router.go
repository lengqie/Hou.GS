package router

import (
    "gin/service"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func InitRouter()  {
    router := gin.Default()
    router.Use(cors.Default())

    // 跳转接口
    router.GET("/:key", func(context *gin.Context) {
        key := context.Param("key")
        status, msg := service.GetUrl(key)
        context.JSONP(http.StatusOK, gin.H{
            "status": status,
            "msg": msg,
        })
    })

    // v1
    v1 := router.Group("v1")
    {
        v1.GET("/url/*url", func(context *gin.Context) {
            url := context.Param("url")
            url = strings.Trim(url, "/")
            //fmt.Println(url)
            // 传入业务层
            status, msg := service.V1Service(url)

            context.JSONP(http.StatusOK, gin.H{
                "status": status,
                "msg": msg,
            })

        })
    }

    // v2
    v2 := router.Group("v2")
    {
        v2.GET("/url/*url", func(context *gin.Context) {
            url := context.Param("url")
            url = strings.Trim(url, "/")
        })
    }

    router.Run(":8080")
}