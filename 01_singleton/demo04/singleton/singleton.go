package singleton

import (
	"sync"
	"sync/atomic"
)

// 单例设计模式 atomic原子实现
// 原子实现确保在多个 goroutine 同时获取实例时不会产生并发安全问题

// singleton指针全局变量
var instance *singleton

// mutex互斥锁全局变量
var mutex sync.Mutex

// done 标志，用于标记instance实例是否初始化
var done uint32

// singleton 单例结构体
type singleton struct {
}

// GetInstance 获取singleton指针实例全局函数
func GetInstance() *singleton {
	if atomic.LoadUint32(&done) == 0 {
		mutex.Lock()         // done 标记为0 上锁
		defer mutex.Unlock() // 最后释放锁
		if done == 0 {
			// 初始化实例
			instance = &singleton{}
			// 修改标记done为1
			atomic.StoreUint32(&done, 1)
		}
	}
	return instance
}
