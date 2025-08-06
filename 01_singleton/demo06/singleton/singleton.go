package singleton

// 单例设计模式 init函数实现

// singleton指针全局变量
var instance *singleton

// singleton 单例结构体
type singleton struct {
}

// init 初始化单例实例
func init() {
	instance = &singleton{}
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	return instance
}
