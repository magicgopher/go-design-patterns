package mediator

import "testing"

// 单元测试
// 模拟客户端调用

// TestMediator 测试中介者模式的场景
func TestMediator(t *testing.T) {
	mediator := NewConcreteChatMediator()

	// 创建用户
	user1 := NewConcreteUser("Alice", mediator)
	user2 := NewConcreteUser("Bob", mediator)
	user3 := NewConcreteUser("Charlie", mediator)

	// 添加用户到聊天室
	mediator.AddUser(user1)
	mediator.AddUser(user2)
	mediator.AddUser(user3)

	// 用户发送消息，通过中介者进行通信
	t.Run("Alice sends a message", func(t *testing.T) {
		t.Log("--- Alice 发送消息 ---")
		user1.Send("大家好！")
		// 预期：Bob 和 Charlie 会收到消息
	})

	t.Run("Bob sends a message", func(t *testing.T) {
		t.Log("\n--- Bob 发送消息 ---")
		user2.Send("你好, Alice！")
		// 预期：Alice 和 Charlie 会收到消息
	})
}
