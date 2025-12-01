package observer

import "fmt"

// 观察者模式（观察者接口与具体实现）

// Observer 定义观察者接口
// 所有具体的观察者都必须实现这个方法，当主题状态改变时会被调用
type Observer interface {
	// Update 当主题状态发生变化时，主题会调用此方法通知观察者
	// 参数 message 是主题传递过来的最新状态信息
	Update(message string)
}

// ConcreteObserver 具体观察者实现
type ConcreteObserver struct {
	name string // 观察者名称，用于区分不同的观察者
}

// NewConcreteObserver 创建一个具体的观察者实例
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

// Update 实现 Observer 接口，接收主题的通知并打印
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("[观察者: %s] 收到状态更新通知 → %s\n", o.name, message)
}
