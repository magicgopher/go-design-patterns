package mediator

// 中介者模式

// 同事类

// User 是同事类接口，定义用户行为
type User interface {
	Send(message string)
	Receive(message string)
}
