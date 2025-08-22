package singleton

// 单例设计模式 无锁实现
// 多个 goroutine 获取实例时会有并发安全问题

// singleton指针全局变量
var instance *singleton

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
