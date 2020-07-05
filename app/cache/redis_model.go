package cache

/*type RedisInterface interface {
	Identify() string                                                 // Redis库标识(此方法可重构,用于切换库,默认master)
	Init()                                                            // Model初始化方法(如有自定义函数,可用此函数初始化Method Map)
	Self(model RedisInterface) RedisInterface                         // 设置自身指针(此方法禁止重构)
	Pool() *redis.Pool                                                // 获取Redis连接池对象(此方法禁止重构)
	Connect() redis.Conn                                              // 获取一个Redis连接(此方法禁止重构)
	Call(methodName string, args ...interface{}) interface{}          // 调用自定义方法(此方法禁止重构)
	Command(name string, args ...interface{}) (*base.AnyValue, error) // 执行redis命令(此方法禁止重构)
}*/