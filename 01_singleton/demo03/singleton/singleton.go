package singleton

import "sync"

// 单例设计模式 双重检查实现
// 判断实例是否为空，才执行加锁操作

// singleton指针全局变量
var instance *singleton

// mutex互斥锁全局变量
var mutex sync.Mutex

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	if instance == nil { // 这里 goroutine 同时获取实例还是有并发安全问题
		mutex.Lock()         // 上锁操作
		defer mutex.Unlock() // 最后释放锁
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}
