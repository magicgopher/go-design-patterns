package singleton

import "sync"

// 单粒设计模式 加锁实现
// 解决多个 goroutine 获取单例实例时并发安全问题
// 获取单例实例会有加锁操作，会有损程序性能

// singleton指针全局变量
var instance *singleton

// mutex互斥锁全局变量
var mutex sync.Mutex

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	mutex.Lock()         // 上锁操作
	defer mutex.Unlock() // 最后释放锁
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
