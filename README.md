# regin框架 #
regin是一款基于go-gin框架封装的web框架,用于快速构建web应用和后端服务.
### 目录
- [安装与配置](https://github.com/go-touch/regin)  
- 快速开始
- 项目结构
- 路由配置

### 安装与配置  
#### 1. 安装Go (version 1.10+), 使用下面命令进行安装regin
	$ go get github.com/go-touch/regin
#### 2. 设置环境变量      
	REGIN_RUNMODE = dev | test | prod
#### 3. 如使用go mod包依赖管理工具,请参考下面命令
##### Windows 下开启 GO111MODULE 并设置 GOPROXY 的命令为：
	$ set GO111MODULE = on
	$ go env -w GOPROXY = https://goproxy.cn,direct
##### MacOS 或者 Linux 下开启 GO111MODULE 并设置 GOPROXY 的命令为：
	$ export GO111MODULE = on
	$ export GOPROXY = https://goproxy.cn
### 快速开始
####入口文件
	$ cat main.go

	package main
	
	import (
		"github.com/go-touch/regin"
		_ "xxx/application/router"
	)
	
	func main() {
		regin.Guide.HttpService()
	}

### 项目结构
- application
    - modules // action存放目录,
  	    - demo // 模块名
    - model
        - bussiness // 业务类
        - common // 公共处理函数
        - dao // 抽象数据模型
        - form // 校验数据模型
        - mysql // mysql表字段映射模型
        - redis // redis数据模型
    - router // 注册路由, 主要通过map存储modules对应的action实例的映射
- config
    - dev // 必要配置项(通过系统环境变量选择配置路径)
        - database.ini
        - server.ini
    - test
    - prod
- runtime
    - log // 存储系统错误日志
- main.go // 入口文件

