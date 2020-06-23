package base

import (
	"github.com/go-touch/regin/utils"
	"errors"
)

/**************************************** 数据类型 - 结构体RouterStorage ****************************************/
// 路由存储器
type RouterStorage struct {
	mode       string     // 路由模式: General、Restful
	generalMap GeneralMap // 对应路由模式: General
}

// 定义Router
var Router *RouterStorage

func init() {
	Router = &RouterStorage{
		mode:       "General",
		generalMap: GeneralMap{},
	}
}

// 设置模式
func (r *RouterStorage) setMode(mode string) error {
	r.mode = mode
	return nil
}

// 获取模式
func (r *RouterStorage) GetMode() string {
	return r.mode
}

// 普通模式
func (r *RouterStorage) General(moduleName string, generalMap GeneralMap) error {
	// 设置模式
	if r.mode != "General" {
		r.setMode("General")
	}
	// 存储
	var moduleKey string
	for key, value := range generalMap {
		moduleKey = utils.StringJoinByDot(moduleName, key)
		r.generalMap[moduleKey] = value
	}
	return nil
}

// 获取GeneralMap
func (r *RouterStorage) GetGeneralMap() GeneralMap {
	return r.generalMap
}

// 获取元素
func (r *RouterStorage) GetGeneral(key string) (value AppAction, err error) {
	if action, ok := r.generalMap[key]; ok {
		value = action
		err = nil
	} else {
		err = errors.New("获取app.Action失败")
	}
	return value, err
}
