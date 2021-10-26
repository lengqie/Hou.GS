package service

import (
    "gin/database"
    "gin/model"
    "regexp"
    "strconv"
    "strings"
)

// 同时实现 service 与 mapper (暂时， 2021年10月24日)
func V1Service(url string) (int, string){

    // 判断url是否合法
    count := strings.Index(url, "http")
    if count != 0 {
        return 1, "请使用http或者https开头"
    }
    reg1 := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
    result1 := reg1.FindAllStringSubmatch(url, -1)
    if len(result1) ==0{
        return 1, "请输入正确的url"
    }

    // 判断url是否存在
    var tempUrl model.Url
    // 初始化数据库
    database.InitDb()
    database.DB.Where("original_url = ?",url).First(&tempUrl)
    // 已经存在
    if(tempUrl != model.Url{}){
        return 0, strconv.Itoa(int(tempUrl.ID))
    }

    // 插入数据库 返回key
    tempUrl = model.Url{
        OriginalUrl: url,
    }
    database.DB.Create(&tempUrl)
    return 0, strconv.Itoa(int(tempUrl.ID))
}

func GetUrl(key string) (int, string) {

    var tempUrl model.Url
    // 初始化数据库
    database.InitDb()
    // 目前用Id替换 Key (暂时，2021年10月25日)
    database.DB.Where("id = ?",key).First(&tempUrl)
    if(tempUrl != model.Url{}){
        return 0, tempUrl.OriginalUrl
    } else {
        return 1, ""
    }

}