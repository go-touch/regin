# regin框架 #
regin是一款基于go-gin框架封装的web框架,用于快速构建web应用和后端服务.
### 目录
- [安装与配置](#安装与配置)
- [快速开始](#快速开始)
- [项目结构](#项目结构)
- [路由配置](#路由配置)
- [服务配置](#服务配置)
- [Web应用](#Web应用)
- [数据库](#数据库)
- [Redis](#Redis)
- [Utils工具](#Utils工具)
	- config文件解析器(支持json、ini)
	- 加密算法
	- 类型处理(自定义类型、基本类型转换)
	- 验证器
	- curl
	- 文件操作
	- ...

### <a id="安装与配置">安装与配置</a>
#### 1. 安装Go (version 1.10+), 然后可使用下面命令进行安装regin
	$ go get github.com/go-touch/regin
#### 2. 设置系统环境变量
	REGIN_RUNMODE = dev | test | prod
#### 3. 依赖包安装(已安装可忽略)
	$ go get github.com/gin-gonic/gin
	$ go get github.com/unrolled/secure
	$ go get github.com/go-sql-driver/mysql
	$ go get github.com/garyburd/redigo/redis
	$ go get gopkg.in/ini.v1
#### 4. 如使用go mod包依赖管理工具, 请参考下面命令
##### Windows 下开启 GO111MODULE 并设置 GOPROXY 的命令为：
	$ set GO111MODULE=on
	$ go env -w GOPROXY=https://goproxy.cn,direct
##### MacOS 或者 Linux 下开启 GO111MODULE 并设置 GOPROXY 的命令为：
	$ export GO111MODULE=on
	$ export GOPROXY=https://goproxy.cn
### <a id="快速开始">快速开始</a>
#### 入口文件: xxx/main.go (xxx 代表项目名称,后面也是)
```go
package main

import (
	"github.com/go-touch/regin"
	_ "xxx/application/router"
)

func main() {
	regin.Guide.HttpService()
}
```
	$ go run main.go
	或者:
	$ go build main.go
	$ ./xxx
### <a id="项目结构">项目结构</a>
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
        - redis.ini
    - test
    - prod
- runtime
    - log // 存储系统错误日志
- main.go // 入口文件

### <a id="路由配置">路由配置</a>
#### 举例: xxx/application/router/demo.go
```go
package router

import (
	"github.com/go-touch/regin/base"
	"mygin/application/modules/api/role"
	"mygin/application/modules/demo"
)

func init() {
	base.Router.General("demo", base.GeneralMap{
		"v1.index":      &demo.Index{}, // demo包下的action：Index
		"role.checkapi": &role.CheckApi{}, // role包下的action：CheckApi
	})
}
```
```go
访问:
http://127.0.0.1/module/controller/action

示例:
http://127.0.0.1/demo/v1/index
http://127.0.0.1/demo/role/checkapi
```
```go
Note: regin中的路由比较松散,url中pathinfo采用三段路径, 通过获取三段路由信息,使用 . 拼接作为key,读取路由map里面对应的action(action定义可查看web应用介绍). 因此路径的含义可依据路由配置定义,并无严格规定.
```
### <a id="路由配置">服务配置</a>
#### 配置项 xxx/config/server.ini
	; 主配置
	[main]
	httpHost = 127.0.0.1:8080 // http服务地址端口
	httpsHost = 127.0.0.1:443 // https服务地址端口

	; 错误配置
	[error]
	log = true // 是否开启错误日志
	读取format = {date}-Error // 日志文件名格式
	pattern = local // 日志存储方式 local:本地 remote:远程
#### 项目中使用配置
	读取方式： (不会直接获取对应值,而是返回一个regin的configValue结构体指针，可实现对应类型转换)
	config ：= service.App.GetConfig("xxx.xxx.xxx")

	调用示例:
	config := service.App.GetConfig("server.main.httpHost").ToString()

	备注: server.ini、database.ini、redis.ini等必要配置项,文件和字段名均为regin使用,不可修改.
### <a id="Web应用">Web应用</a>
#### 基类action的代码示例: regin/base/app.action.go
```go
package base

// action接口
type AppAction interface {
	BeforeExec(request *Request) (result *Result) // 调用方法 Exec 前执行, 可通过其实现token验证、鉴权等业务
	Exec(request *Request) (result *Result) // 用于实现具体的业务逻辑
}

// action接口的一个实现
type Action struct {
	AppAction
}

// Before action method
func (a *Action) BeforeExec(request *Request) (result *Result) {
	return
}

// Action method.
func (a *Action) Exec(request *Request) (result *Result) {
	return
}
```
#### 项目action的代码示例: xxx/application/modules/demo/mysql_select.go
```go
package demo

import (
	"github.com/go-touch/regin/app/db"
	"github.com/go-touch/regin/base"
)

type MysqlSelect struct {
	base.Action // 继承基类action
}

// 执行方法(重载方法)
func (this *MysqlSelect) Exec(request *base.Request) *base.Result {
	result := base.JsonResult()

	// 查询一条数据
	ret := db.Model("PlusArticle").FetchRow(func(dao *db.Dao) {
		dao.Where("id", 202)
	})

	// 错误处理
	if err := ret.ToError(); err != nil {
		result.SetData("code", 1000)
		result.SetData("msg", "系统出了点问题~")
		return result
	}

	// json输出数据库查询结果
	result.SetData("data", ret.ToStringMap())
	return result
}
```
#### *base.Request实例(封装param、get、post方法,自动json、xml解析)
##### 获取请求方式
	GetMethod() string
##### 获取error信息
	GetError() error
##### 获取pathinfo的路径信息. 参数key:路径名 defaultValue:默认值
	Param(key string, defaultValue ...string)
##### 获取pathinfo的路径map信息. 返回值为 base.StringMap(对应基础类型map[string]string)
	ParamAll() StringMap
##### 获取Post数据, json、xml数据也用此方法.返回值为 interface{}
	Post(key string, defaultValue ...interface{}) (value interface{}, err error)
##### 获取Post数据的map, json、xml数据也用此方法. 返回值为 base.AnyMap(对应基础类型map[string]interface{})
	PostAll() (anyMap AnyMap, err error) 获取一个map[string]interface{}
##### 获取上传文件io句柄. 返回值为 []*multipart.FileHeader
	PostFile(name string) []*multipart.FileHeader
##### 更多方法使用 request 调用
	...
#### *base.Result实例(用于响应客户端)
```go
type Result struct {
	Type   string // 可选值为:String、Json、Html、
	Page   string // 响应页面(Type = Html时必填)
	Status int    // 状态码 200正常状态
	Msg    string // 提示消息
	Data   AnyMap // 业务数据
}

// 定义RespResult
var ResultInvoker *Result

func init() {
	ResultInvoker = &Result{}
}

// 创建Json result
func (r *Result) CreateJson(status int, msg string) *Result {
	return &Result{
		Type:   "Json",
		Page:   "",
		Status: status,
		Msg:    msg,
		Data:   AnyMap{"code": 0, "msg": "", "data": ""},
	}
}

// 创建Html result
func (r *Result) CreateHtml(page string, status int, msg string) *Result {
	return &Result{
		Type:   "Html",
		Page:   page,
		Status: status,
		Msg:    msg,
		Data:   AnyMap{},
	}
}
```
##### 获取一个可响应json的 *base.Result 实例
	base.ResultInvoker.CreateJson(status int, msg string) *Result
##### 获取一个可响应html的 *base.Result 实例
	base.ResultInvoker.CreateHtml(status int, msg string) *Result
##### 快速获取一个可响应json的  *base.Result 实例
	base.JsonResult() *Result
##### 修改业务数据即 *base.Result 的 Data 字段
	(r *Result) SetData(key string, value interface{})
##### 获取业务数据即 *base.Result 的 Data 字段
	(r *Result) GetData(key string) interface{} 
#### *base.AnyValue值类型（用于数据转换,对于不确定类型interfa{}比较适用)
##### 获取 *base.AnyValue. 参数value:interface{}(可传任意值)
	base.Eval(value interface{}) *AnyValue
##### 返回错误信息
	(av *AnyValue) ToError() error
##### 返回原值
	(av *AnyValue) ToValue() interface{}
##### 转成int类型
	(av *AnyValue) ToInt() int
##### 转成byte类型
	(av *AnyValue) ToByte() byte
##### 转成string类型
	(av *AnyValue) ToString() string
##### 转成bool类型
	(av *AnyValue) ToBool() bool
##### 转成map[string]string类型
	(av *AnyValue) ToStringMap() map[string]string
##### 更多方法使用 *base.AnyValue 调用
	...
#### regin定义的数据类型 (业务中可直接使用,包名base)
```go
// 预定义常见数据类型
type DataType interface {
	Set(key string, value interface{})
	Get(key string) *AnyValue
}
type AnyMap map[string]interface{}        // [MapType] key is string,value is 任意类型
type StringMap map[string]string          // [MapType] key is string,value is string 类型
type IntMap map[string]int                // [MapType] key is string,value is int 类型
type StringSliceMap map[string][]string   // [MapType] key is string,value is string Slice 类型
type GeneralMap map[string]AppAction      // [MapType] key is string,value is AppAction t类型
type AnySlice []interface{}               // [SliceType] key is index,value为任意类型
type StringMapSlice []map[string]string   // [SliceType] key is index,value为(key为string,value为string)的map
type AnyMapSlice []map[string]interface{} // [SliceType] key is index,value为(key为string,value为任意类型)的map
```
```go
Note: 部分值为 interface{} 的类型实现了 DataType 接口, 需要类型转换可通过Get方法获取到一个 *base.AnyValue
```
### <a id="数据库">数据库</a>
#### 配置项 xxx/config/dev/database.ini
	[plus_center] // 配置分组,必填
	; 主库
	master.driverName = mysql // 驱动名称
	master.dataSourceName = root:root@tcp(127.0.0.1:3306)/dbName?charset=utf8 // 连接参数
	master.maxIdleConn = 100 // 空闲连接数
	master.maxOpenConn = 100 // 最大连接数

	; 从库
	slave.driverName = mysql
	slave.dataSourceName = root:root@tcp(127.0.0.1:3306)/dbName?charset=utf8
	slave.maxIdleConn = 100
	slave.maxOpenConn = 100
#### Model的示例
```go
package mysql

import (
	"github.com/go-touch/regin/app/db"
)

type Users struct{
	Id       string `field:id`
	Username string `field:"username"`
}

// 注册model
func init() {
	db.RegisterModel(&Users{}, "Users")
}

// 数据库标识(此方法可重构,用于切换数据库,默认master)
func (this *Users) Identify() string {
	return "plus_center.master"
}

// 数据库表名(此方法可重构,用于切换数据表)
func (this *Users) TableName() string {
	return "users"
}

// 自定义方法
func (this *Users) Method() string {
	ret := db.Model("Users").FetchAll(func(dao *db.Dao) {
		dao.Where("id", 202)
	})
}
```
```go
(this *Users) Identify() string //设置数据库连接标识,对应数据库配置的key链关系
```
```go
(this *Users) TableName() string // 设置真实数据表名,如未设置则默认结构体名称(注: AdminUser 会转成 admin_user)
```
#### 使用Model查询一条记录示例:
```go
第一种方式:
row := db.Model(&Users{}).FetchRow(func(dao *db.Dao) {
	dao.Where("id", 1)
})

第二种方式:
// 注册 Model, 第一个参数传入model实例化的指针, 第二个可选参数,用于起别名方便调用,不传名称默认为 mysql.Users .
db.RegisterModel(&Users{}, "Users")

// 使用别名获取 Dao 数据对象并使用 FetchRow 方法查询
row := db.Model("Users").FetchRow(func(dao *db.Dao) {
	dao.Where("id", 1)
})
```
```go
Node: 推荐使用第二种方式,可以在初始化函数 init 批量注册model,这样在系统加载的时候仅调用一次注入容器.
```
#### db.Dao方法(举例均采用上述的第二种方式)
##### 获取Dao数据对象.
```go
db.Model(userModel interface{})

调用示例:
dao := db.Model(&Users{})
或
db.RegisterModel(&Users{}, "Users")
dao := db.Model("Users")
```
##### 设置表名.(通常无需调用,注册model时已获取表名)
```go
(d *Dao) Table(tableName string) *Dao

示例:
db.Model("Users").Table("message")
```
##### 设置表字段.参数field: 可为string或[]string
```go
(d *Dao) Field(field interface{}) *Dao

示例:
db.Model("Users").Field("a,b,c,d")
db.Model("Users").Field([]string{"a,b,c,d"})
```
##### 设置查询条件. 参数field:字段名,参数value:字段值,参数linkSymbol:连接符(and[or] 默认and)
```go
(d *Dao) Where(field interface{}, value interface{}, linkSymbol ...string) *Dao

示例:
db.Model("Users").Where("id", 1)
```
##### 批量设置查询条件,和where类似. 参数是 key-value 的 map
```go
(d *Dao) WhereMap(fieldMap map[string]interface{}, linkSymbol ...string) *Dao

示例:
db.Model("Users").WhereMap(map[string]interface{}{"id":1})
```
##### 绑定数据,insert[update]时使用到.
```go
(d *Dao) Values(valueMap map[string]interface{}) *Dao

示例:
db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"})
```
##### 设置排序. 参数不定
```go
(d *Dao) Order(expr ...string) *Dao

示例:
db.Model("Users").Order("id ASC","username Desc")
```
##### 批量设置排序. 与Order类似,参数为[]string
```go
(d *Dao) OrderSlice(expr []string) *Dao

示例:
db.Model("Users").OrderSlice([]string{"id ASC","username Desc"})
```
##### 设置查询检索记录行. 参数不定,对应sql语句 limit m,n
```go
(d *Dao) Limit(limit ...int) *Dao

示例:
db.Model("Users").Limit(1,10)
```
##### 是否返回sql. 需要在执行增删改查方法前调用
```go
(d *Dao) Sql() *Dao

示例:
ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
	dao.Sql()
})
fmt.Println(ret.ToString()) // 打印字符串sql语句
```
##### 查询一条记录,返回 \*db.AnyValue,可实现数据转换.参数userFunc:用户回调函数,接收参数为 \*db.Dao
```go
(d *Dao) FetchRow(userFunc ...UserFunc) *AnyValue

示例:
ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
	dao.Where("id", 1)
})
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToStringMap // 返回 map[string]string 结构的一条数据
```
##### 查询多条记录,返回 \*db.AnyValue,可实现数据转换. 参数userFunc:用户回调函数,接收参数为 \*db.Dao
```go
(d *Dao) FetchAll(userFunc ...UserFunc) *AnyValue

示例:
ret := db.Model("Users").FetchAll(func(dao *db.Dao) {
	dao.Where("id", 1)
})
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToStringMap // 返回 map[string]string 结构的一条数据
```
##### 插入一条记录,返回 \*db.AnyValue,可实现数据转换. 参数userFunc:用户回调函数,接收参数为 \*db.Dao
```go
(d *Dao) Insert(userFunc ...UserFunc) *AnyValue

示例:
ret := db.Model("Users").Insert(func(dao *db.Dao) {
	dao.Values(map[string]interface{}{"username":"zhangsan"})
})
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToLastInsertId() // 返回最后插入的主键id
```
##### 更新一条记录,返回 \*db.AnyValue,可实现数据转换. 参数userFunc:用户回调函数,接收参数为 \*db.Dao
```go
(d *Dao) Update(userFunc ...UserFunc) *AnyValue

示例:
ret := db.Model("Users").Update(func(dao *db.Dao) {
	dao.Values(map[string]interface{}{"username":"zhangsan"})
})
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToAffectedRows() // 返回受影响行数
```
#####  删除一条数据,返回 \*db.AnyValue,可实现数据转换. 参数userFunc:用户回调函数,接收参数为 \*db.Dao
```go
(d *Dao) DELETE(userFunc ...UserFunc) *AnyValue

示例:
ret := db.Model("Users").DELETE(func(dao *db.Dao) {
	dao.Where("id", 1)
})
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToAffectedRows() // 返回受影响行数
```
#### 完整数据库操作实例(假设model已注册)
##### 查询一条数据
```go
// 链式操作
ret := db.Model("Users").Field("username").Where("id", 1).FetchRow()

// 匿名函数回调操作
ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
	dao.Field("username")
	dao.Where("id", 1)
})
ret.ToError()
ret.ToStringMap()
```
##### 查询多条数据
```go
// 链式操作
ret := db.Model("Users").Field("username").Where("id", 1).FetchAll()

// 匿名函数回调操作
ret := db.Model("Users").FetchAll(func(dao *db.Dao) {
	dao.Field("username")
	dao.Where("id", 1)
})
ret.ToError()
ret.ToStringMapSlice()
```
##### 插入一条数据
```go
// 链式操作
ret := db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"}).Insert()

// 匿名函数回调操作
ret := db.Model("Users").Insert(func(dao *db.Dao) {
	dao.Values(map[string]interface{}{"username":"zhangsan"})
})
ret.ToError()
ret.ToLastInsertId()
```
##### 更新一条数据
```go
// 链式操作
ret := db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"}).Where("id", 1).Update()

// 匿名函数回调操作
ret := db.Model("Users").Update(func(dao *db.Dao) {
	dao.Values(map[string]interface{}{"username":"zhangsan"})
	dao.Where("id", 1)
})
ret.ToError()
ret.ToAffectedRows()
```
##### 删除一条数据
```go
// 链式操作
ret := db.Model("Users").Where("id", 1).DELETE()

// 匿名函数回调操作
ret := db.Model("Users").DELETE(func(dao *db.Dao) {
	dao.Where("id", 1)
})
ret.ToError()
ret.ToAffectedRows()
```
### <a id="Redis">Redis</a>

#### 配置项 xxx/config/dev/redis.ini
	[plus_center] // // 配置分组,必填
	master.host = 127.0.0.1:6379 // 主机端口
	master.password = "" // 密码
	master.db = 10 // 库标
	master.MaxIdle = 16 // 空闲连接数
	master.MaxActive = 32 // 最大连接数 
	master.IdleTimeout = 120 // 超时时间

#### RedisModel的示例
```go
package redis

type TestModel struct {}

// Redis库标识
func (b *Base) Identify() string {
	return "plus_center.master"
}
```
```go
(this *Users) Identify() string // 设置redis连接参数,对应Redis配置的key链关系
```
#### Redis的使用示例:
##### 传入一个model获取 Redis Dao 实例
```go
RedisModel(model interface{}) *RedisDao

示例:
redisDao := RedisModel(&TestModel{})
```
##### 获取连接池对象,开发者可通过此返回值
```go
(rd *RedisDao) Pool() *redis.Pool

示例:
pool := RedisModel(&TestModel{}).Pool()
```
##### 执行redis命令,返回\*base.AnyValue,可进行类型转换. 参数name:命令名称 args:该命令对应的参数
```go
(rd *RedisDao) Command(name string, args ...interface{}) *base.AnyValue

示例:
RedisModel(&TestModel{}).Command("SET","username","admin")
RedisModel(&TestModel{}).Command("HSET","user","username","admin")
ret := RedisModel(&TestModel{}).Command("GET","username")
ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
ret.ToAffectedRows() // 返回受影响行数
```
### <a id="Utils工具">Utils工具</a>
#### Form表单验证
##### Form Model结构体示例
```go
type PlusUsers struct {
	UserId  int    `key:"user_id" require:"true" length:"0|5"`
	Account string `key:"account" require:"true" length:"0|20"`
}
```
##### Form验证器的使用
```go
第一种方式:
result := validator.Form(&PlusUsers{}).Verify(&map[string]interface{}{
	"user_id": 1,
})

第二种方式:
validator.RegisterForm(&PlusUsers{}, "PlusUsers")
result := validator.Form("PlusUsers").Verify(&map[string]interface{}{
	"user_id": 1,
})
```
##### Form验证器方法:
##### 获取一个Form Dao
```go
// 获取 Form Dao
Form(userModel interface{}) *FormHandle

示例:
formDao :=  validator.Form(&PlusUsers{})
```
##### 获取一个Form Dao(另一种方式)
```go
// 注册 Form Model
RegisterForm(userModel interface{}, alias ...string)

// 获取 Form Dao
Form(userModel interface{}) *FormHandle

示例:
validator.RegisterForm(&PlusUsers{}, "PlusUsers")
formDao := validator.Form("PlusUsers")
```
##### 验证一个*map[string]interface{}
```go
(mh *FormHandle) Verify(vMap *map[string]interface{}) []*tag.Result
```


