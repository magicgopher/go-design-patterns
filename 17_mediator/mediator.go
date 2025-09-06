package mediator

// 中介者模式

// 中介者接口

// ChatMediator 定义聊天室中介者接口
type ChatMediator interface {
	SendMessage(message string, sender User)
	AddUser(user User)
}
