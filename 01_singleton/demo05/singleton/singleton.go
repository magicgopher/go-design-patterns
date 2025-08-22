package singleton

import "sync"

// 单例设计模式 sync.Once 实现

// singleton指针全局变量
var instance *singleton

// 定义once全局变量
var once sync.Once

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	// once.Do() 保证变量仅被初始化一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
