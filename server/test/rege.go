package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {

    buf := "://baidu.com"

    //解析正则表达式，如果成功返回解释器
    reg1 := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
    //根据规则提取关键信息
    result1 := reg1.FindAllStringSubmatch(buf, -1)
    //fmt.Println("result1 = ", len(result1[0]))
    fmt.Println("result1 = ", result1)
    fmt.Println(strings.Index("http://baidu.com","h"))
}