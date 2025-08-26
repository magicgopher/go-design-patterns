package bridge

import "fmt"

// 桥接模式 - 实现部分 (Implementor)

// MessageSender 定义了发送消息的接口
// 这是“实现”维度的抽象
type MessageSender interface {
	Send(message string)
}

// 具体实现部分 (Concrete Implementor)

// SmsSender 短信发送器
type SmsSender struct{}

// NewSmsSender 创建SmsSender实例
func NewSmsSender() *SmsSender {
	return &SmsSender{}
}

// Send 实现MessageSender接口，通过短信发送消息
func (s *SmsSender) Send(message string) {
	fmt.Printf("通过【短信】发送消息: %s\n", message)
}

// EmailSender 邮件发送器
type EmailSender struct{}

// NewEmailSender 创建EmailSender实例
func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

// Send 实现MessageSender接口，通过邮件发送消息
func (e *EmailSender) Send(message string) {
	fmt.Printf("通过【邮件】发送消息: %s\n", message)
}
