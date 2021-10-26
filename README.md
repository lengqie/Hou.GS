<div align=center>
<img src="./static/hou.png"/>
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.14.2-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.6-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.21.16-red"/>
<img src="https://img.shields.io/badge/vue-3.2.16-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-1.1.0beta24-green"/>
</div>


## 1. 基本介绍

### 1.1 简介

一个基于Gin、Vue，前后端分离的短网址项目。

![](.\static\dome.jpg)

### 1.2 功能

- [x]   基本功能

- [ ]   部署

- [ ]   后端key加密

- [ ]   移动端适应

- [ ]   中转界面定义

- [ ]   后台

- [ ]   Docker部署

    

非礼勿视，非礼勿听，非礼勿言

## 2. 使用说明

### 2.1 server项目

使用 `Goland` 等编辑工具，打开server目录

```bash
# 克隆项目
git clone https://github.com/lengqie/Hou.GS.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate
```
```
# 数据库迁移
#修改数据库信息 \database\mysql.go 第20行
#运行 main.go 方法中的database.Migrate()方法
```

```bash
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
cnpm install || npm install

# 启动web项目
npm run serve
```

### 2.3部署上线


- server

  ```bash
  cd server
  
  # 使用 go mod 并安装go依赖包
  go generate
  
  # 编译 
  go build -o server main.go (windows编译命令为go build -o server.exe main.go )
  
  # 运行二进制
  ./server (windows运行命令为 server.exe)
  ```
  
  ```
  # ningx
  ```


- web

  ```bash
  # 进入web文件夹
  cd web
  
  # 安装依赖
  cnpm install || npm install
  ```
  
  ```
  # 修改api地址
  #修改\web\src\components\Main.vue 第99行
  #修改\web\src\components\Redirect.vue 第14行
  ```
  
  ```bash
  # build
  npm run build
  ```
  


## 3. 协议

Licensed under the *MIT* license