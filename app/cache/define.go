package cache

// 构建Redis Model
func CreateRedis(model RedisInterface) RedisInterface{
	model.Self(model)
	model.Init()
	return model
}