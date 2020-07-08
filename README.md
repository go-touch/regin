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

### <a id="安装与配置">安装与配置</a>
#### 1. 安装Go (version 1.10+), 然后可使用下面命令进行安装regin
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
### <a id="快速开始">快速开始</a>
####入口文件(xxx 代表项目名称,后面也是)
	$ cat ./xxx/main.go

	package main
	
	import (
		"github.com/go-touch/regin"
		_ "xxx/application/router"
	)
	
	func main() {
		regin.Guide.HttpService()
	}

	$ go run main.go

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
	package router

	import (
		"github.com/go-touch/regin/base"
		"xxx/application/modules/demo"
		"xxx/application/modules/api/role"
	)
	
	func init() {
		base.Router.General("demo", base.GeneralMap{
			"v1.index":     &demo.Index{}, // 对应 modules下的index结构体
			"role.checkapi": &role.CheckApi{}, // 对应 modules下的checkapi结构体
		})
	}

#### 路径访问
	格式: http://127.0.0.1/module/controller/action
	例如: http://127.0.0.1/demo/v1/index
	备注: regin中的路由比较松散,url中pathinfo采用三段路径, 通过获取三段路由信息,使用 . 拼接作为key,读取路由map里面对应的action, 
	(action定义可查看web应用介绍) 因此路径的含义可依据路由配置定义,并无严格规定.
### <a id="路由配置">服务配置</a>
#### xxx/config/server.ini
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

	备注: server.ini、database.ini、redis.ini等必要配置项,文件字段名均为regin使用,不可修改.
		
### <a id="Web应用">Web应用</a>
#### 基类action的代码示例:
	package base
	
	type AppAction interface {
		BeforeExec(request *Request) (result *Result)
		Exec(request *Request) (result *Result)
	}
	
	type Action struct {
		AppAction
	}
	
	// Before action method. // 需实现该方法,在调用方法 Exec 前执行, 可通过其实现token验证、鉴权等业务,通过返回值*result,控制响应结果
	func (a *Action) BeforeExec(request *Request) (result *Result) { 
		return
	}
	
	// Action method. // 需实现该方法,用于实现具体的业务逻辑
	func (a *Action) Exec(request *Request) (result *Result) { 
		return
	}
#### 项目action的代码示例:	
	$ cat xxx/application/modules/demo/mysql_select
	package demo
	
	import (
		"github.com/go-touch/regin/app/db"
		"github.com/go-touch/regin/base"
	)
	
	type MysqlSelect struct {
		base.Action
	}
	
	// 执行方法
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

#### *base.Request实例(封装param、get、post方法,自动json、xml解析)
> GetMethod() string 获取请求方式  
> GetError() error 获取error信息  
> Param(key string, defaultValue ...string) 获取pathinfo的路径信息    
> ParamAll() StringMap // 获取一个map[string]string, 类型属于regin的StringMap  
> Post(key string, defaultValue ...interface{}) (value interface{}, err error)     
> PostAll() (anyMap AnyMap, err error) 获取一个map[string]interface{}, 类型属于regin的AnyMap  
> PostFile(name string) []*multipart.FileHeader 用于获取文件io句柄  
> ...

#### *base.Result实例(用于响应客户端)
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
> ResultInvoker.CreateJson() // ResultInvoker为预定义的 *base.Result 实例  
> (r *Result) CreateJson(status int, msg string) *Result // 创建一个可返回json数据的\*base.Result   
> (r *Result) CreateHtml(page string, status int, msg string) *Result // 创建一个可返回html数据的\*base.Result  
> (r *Result) SetData(key string, value interface{}) // 修改业务数据即 \*base.Result 的 Data 字段   
> (r *Result) GetData(key string) interface{} // 获取业务数据即 \*base.Result 的 Data 字段   

#### *AnyValue值类型（用于数据转换,对于不确定类型interfa{}比较适用,包名base)
> Eval(value interface{}) *AnyValue 通过调用此方法获取 *AnyValue  
> (av *AnyValue) ToError() error 返回错误信息  
> (av *AnyValue) ToValue() interface{} 返回原值  
> (av *AnyValue) ToInt() int 转成int类型  
> (av *AnyValue) ToByte() byte 转成byte类型  
> (av *AnyValue) ToString() string 转成string类型  
> (av *AnyValue) ToBool() bool 转成bool类型  
> (av *AnyValue) ToStringMap() map[string]string 转成map[string]string类型  
> ...

#### regin定义的数据类型 (业务中可直接使用,包名base)
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
	
	备注: 部分值为 interface{} 的类型实现了 DataType 接口, 需要类型转换可通过Get方法获取到一个 *AnyValue

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
> (this *Users) Identify() string 返回数据库连接参数,对应数据库配置的key链关系  
> (this *Users) TableName() string 返回真实数据表名,如未设置则默认结构体名称(注: AdminUser 会转成 admin_user)

#### 使用Model查询一条记录示例:
	
	第一种方式:
	row := db.Model(&Users{}).FetchAll(func(dao *db.Dao) {
		dao.Where("id", 1)
	})

	第二种方式:
	db.RegisterModel(&Users{}, "Users") // 注册model, 第一个参数传入model实例化的指针, 第二个可选参数,用于起别名方便调用.
	row := db.Model("Users").FetchAll(func(dao *db.Dao) {
		dao.Where("id", 1)
	})
	
	note: 推荐使用第二种方式,可以在初始化函数 init 批量注册model,这样在系统加载的时候回调用一次注入容器.

#### db.Dao方法(举例均采用上述的第二种方式)

##### <font color=#0099ff>获取Dao数据对象</font>
##### Model(userModel interface{}) *Dao // 获取Dao数据对象
	db.Model(&Users{})
	或
	db.RegisterModel(&Users{}, "Users")
	db.Model("Users")
##### Table(tableName string) *Dao // 设置表名(通常无需调用,注册model时已获取表名) 
	db.Model("Users").Table("message")
##### Field(field interface{}) *Dao // 设置表字段,参数 field可为string或[]string
	db.Model("Users").Field("a,b,c,d")
	db.Model("Users").Field([]string{"a,b,c,d"})
##### Where(field interface{}, value interface{}, linkSymbol ...string) *Dao // 设置查询条件 参数field: 字段名 参数value: 字段值 参数linkSymbol: 连接符 and[or] 默认and
	db.Model("Users").Where("id", 1)
##### WhereMap(fieldMap map[string]interface{}, linkSymbol ...string) *Dao // 和where类似,参数是key-value的map
	db.Model("Users").WhereMap(map[string]interface{}{"id":1})
##### Values(valueMap map[string]interface{}) *Dao // 绑定数据 insert[update]时使用到
	db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"})
##### Order(expr ...string) *Dao // 设置排序 参数不定
	db.Model("Users").Order("id ASC","username Desc")
##### OrderSlice(expr []string) *Dao // 与Order类似 参数为[]string
	db.Model("Users").OrderSlice([]string{"id ASC","username Desc"})
##### Limit(limit ...int) *Dao // 参数不定,对应sql语句 limit m,n
	db.Model("Users").Limit(1,10) 
##### Sql() *Dao // 是否返回sql
	ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
		dao.Sql()
	})
	fmt.Println(ret.ToString()) // 打印字符串sql语句
##### FetchRow(userFunc ...UserFunc) *AnyValue // 查询一条记录,返回\*db.AnyValue,可实现数据转换.参数userFunc:用户回调函数,接收Dao
	ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
		dao.Where("id", 1)
	})
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToStringMap // 返回 map[string]string 结构的一条数据
##### FetchAll(userFunc ...UserFunc) *AnyValue // 查询多条记录,返回*db.AnyValue,可实现数据转换
	ret := db.Model("Users").FetchAll(func(dao *db.Dao) {
		dao.Where("id", 1)
	})
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToStringMap // 返回 map[string]string 结构的一条数据
##### Insert(userFunc ...UserFunc) *AnyValue // 插入一条数据,返回\*db.AnyValue,可实现数据转换
	ret := db.Model("Users").Insert(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{"username":"zhangsan"})
	})
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToLastInsertId() // 返回最后插入的主键id
##### Update(userFunc ...UserFunc) *AnyValue // 更新一条数据,返回\*db.AnyValue,可实现数据转换
	ret := db.Model("Users").Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{"username":"zhangsan"})
	})
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToAffectedRows() // 返回受影响行数
##### DELETE(userFunc ...UserFunc) *AnyValue // 删除一条数据,返回\*db.AnyValue,可实现数据转换
	ret := db.Model("Users").DELETE(func(dao *db.Dao) {
		dao.Where("id", 1)
	})
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToAffectedRows() // 返回受影响行数
#### 完整数据库操作实例(假设model已注册)
##### 查询一条数据
	// 链式操作
	ret := db.Model("Users").Field("username").Where("id", 1).FetchRow()

	// 匿名函数回调操作
	ret := db.Model("Users").FetchRow(func(dao *db.Dao) {
		dao.Field("username")
		dao.Where("id", 1)
	})
	ret.ToError()
	ret.ToStringMap()
##### 查询多条数据
	// 链式操作
	ret := db.Model("Users").Field("username").Where("id", 1).FetchAll()

	// 匿名函数回调操作
	ret := db.Model("Users").FetchAll(func(dao *db.Dao) {
		dao.Field("username")
		dao.Where("id", 1)
	})
	ret.ToError()
	ret.ToStringMapSlice()
##### 插入一条数据
	// 链式操作
	ret := db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"}).Insert()

	// 匿名函数回调操作
	ret := db.Model("Users").Insert(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{"username":"zhangsan"})
	})
	ret.ToError()
	ret.ToLastInsertId()

##### 更新一条数据
	// 链式操作
	ret := db.Model("Users").Values(map[string]interface{}{"username":"zhangsan"}).Where("id", 1).Update()

	// 匿名函数回调操作
	ret := db.Model("Users").Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{"username":"zhangsan"})
		dao.Where("id", 1)
	})
	ret.ToError()
	ret.ToAffectedRows()
##### 删除一条数据
	// 链式操作
	ret := db.Model("Users").Where("id", 1).DELETE()

	// 匿名函数回调操作
	ret := db.Model("Users").DELETE(func(dao *db.Dao) {
		dao.Where("id", 1)
	})
	ret.ToError()
	ret.ToAffectedRows()

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
	package redis

	type TestModel struct {}
	
	// Redis库标识
	func (b *Base) Identify() string {
		return "plus_center.master"
	}
> (this *Users) Identify() string 返回redis连接参数,对应数据库配置的key链关系

#### Redis的使用示例:
##### RedisModel(model interface{}) *RedisDao // 传入一个model获取Dao实例
	RedisModel(&TestModel{})
##### Pool() *redis.Pool // 获取连接池对象,开发者可通过此返回值
	RedisModel(&TestModel{}).Pool()
##### Command(name string, args ...interface{}) *base.AnyValue // 执行redis命令,返回\*base.AnyValue,可进行类型转换. 参数name:命令名称 args:该命令对应的参数
	RedisModel(&TestModel{}).Command("SET","username","admin")
	RedisModel(&TestModel{}).Command("HSET","user","username","admin")
	ret := RedisModel(&TestModel{}).Command("GET","username")
	ret.ToError() // 可获取错误信息,如果返回nil,则说明无错误发生
	ret.ToAffectedRows() // 返回受影响行数

### <a id="Utils工具">Utils工具</a>
1. Curl功能
2. File基本操作
3. Convert数据类型转换
4. ConfigParser配置解析器
5. ...

```javascript
var ihubo = {
  nickName  : "草依山",
  site : "http://jser.me"
}
```