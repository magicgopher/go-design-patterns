package bridge

import "fmt"

// 桥接模式 - 抽象部分 (Abstraction)

// Message 定义了消息的抽象接口
// 这是“抽象”维度的抽象，它持有一个实现部分的引用 (sender)
type Message interface {
	Send(content string)
}

// 具体抽象部分 (Refined Abstraction)

// CommonMessage 普通消息
type CommonMessage struct {
	sender MessageSender // 桥接的核心：持有一个实现部分的接口
}

// NewCommonMessage 创建一个普通消息实例，并“桥接”一个发送器
func NewCommonMessage(sender MessageSender) *CommonMessage {
	return &CommonMessage{sender: sender}
}

// Send 通过桥接的发送器发送消息
func (m *CommonMessage) Send(content string) {
	// 在发送前可以有一些自己的业务逻辑
	m.sender.Send(content)
}

// UrgencyMessage 加急消息
type UrgencyMessage struct {
	sender MessageSender // 桥接的核心：持有一个实现部分的接口
}

// NewUrgencyMessage 创建一个加急消息实例，并“桥接”一个发送器
func NewUrgencyMessage(sender MessageSender) *UrgencyMessage {
	return &UrgencyMessage{sender: sender}
}

// Send 通过桥接的发送器发送消息
func (m *UrgencyMessage) Send(content string) {
	// 加急消息有自己的业务逻辑，比如添加前缀
	urgencyContent := "【加急】" + content
	m.sender.Send(urgencyContent)
	// 还可以增加额外的处理，比如监控
	fmt.Println("加急消息已发送，启动监控流程...")
}
