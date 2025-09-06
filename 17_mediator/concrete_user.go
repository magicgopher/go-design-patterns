package mediator

import "fmt"

// 中介者模式

// 具体同事类

// ConcreteUser 是具体用户，参与聊天室通信
type ConcreteUser struct {
	name     string
	mediator ChatMediator
}

// NewConcreteUser 创建具体用户实例
func NewConcreteUser(name string, mediator ChatMediator) *ConcreteUser {
	return &ConcreteUser{
		name:     name,
		mediator: mediator,
	}
}

// Send 发送消息，通过中介者广播
func (u *ConcreteUser) Send(message string) {
	fmt.Printf("%s 发送消息: %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}

// Receive 接收消息
func (u *ConcreteUser) Receive(message string) {
	fmt.Printf("%s 收到消息: %s\n", u.name, message)
}
