package main

import "gin/router"

func main() {
    // 初始化数据库
    //database.Migrate()
    //fmt.Println("hello world")

    router.InitRouter()
}