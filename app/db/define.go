package db

// 构建Model
func CreateModel(model ModelInterface) ModelInterface{
	model.Self(model)
	model.Init()
	return model
}
