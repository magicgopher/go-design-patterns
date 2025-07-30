package singleton

import "sync"

// 单例设计模式

// 定义全局单例实例变量
var instance *singleton

// 定义once变量
var once sync.Once

// singleton 单例实例结构体
type singleton struct {
}

// GetInstance 提供全局获取单例实例的函数
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
