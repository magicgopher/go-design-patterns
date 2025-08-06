package singleton

import (
	"fmt"
	"sync"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestGetInstance 模拟client端调用单例（单个goroutine）
func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		fmt.Println("instance1不等于instance2")
	}
}

// TestMultiGetInstance 模拟client端调用单例（多个goroutine）
func TestMultiGetInstance(t *testing.T) {
	// 定义并发执行的 goroutine 数量
	parCount := 100
	// 创建一个 sync.WaitGroup 用于等待所有 goroutine 完成
	wg := sync.WaitGroup{}
	// 创建一个存储单例实例的数组，用于保存每个 goroutine 获取的单例对象
	instances := make([]*singleton, parCount)

	// 启动并发 goroutines 来获取单例实例
	for i := 0; i < parCount; i++ {
		wg.Add(1) // 每启动一个 goroutine，都增加一个计数
		go func(index int) {
			defer wg.Done()                  // 每个 goroutine 完成后，调用 Done() 来减少计数
			instances[index] = GetInstance() // 获取单例实例并存储在对应位置
		}(i)
	}

	// 等待所有的 goroutines 完成
	wg.Wait() // 阻塞，直到所有 goroutine 调用 Done() 完成

	// 检查所有 goroutine 中获取的单例实例是否相同
	for i := 1; i < parCount; i++ {
		// 如果某个实例不与第一个实例相同，说明单例实现存在问题
		if instances[i] != instances[0] {
			t.Errorf("预期所有实例都相同，但实例 %d 不同。", i)
		}
	}
}
