package mediator

import "sync"

// 中介者模式

// 具体中介者

// ConcreteChatMediator 是具体中介者，管理用户并协调消息发送
type ConcreteChatMediator struct {
	users []User
	mu    sync.RWMutex
}

// NewConcreteChatMediator 创建具体中介者实例
func NewConcreteChatMediator() *ConcreteChatMediator {
	return &ConcreteChatMediator{
		users: make([]User, 0),
	}
}

// AddUser 添加用户到聊天室 (写入操作，使用写锁)
func (m *ConcreteChatMediator) AddUser(user User) {
	m.mu.Lock()         // 加锁
	defer m.mu.Unlock() // 保证函数结束时解锁
	m.users = append(m.users, user)
}

// SendMessage 广播消息给除发送者外的所有用户 (读取操作，使用读锁)
func (m *ConcreteChatMediator) SendMessage(message string, sender User) {
	m.mu.RLock()         // 加读锁
	defer m.mu.RUnlock() // 保证函数结束时解锁
	for _, user := range m.users {
		if user != sender {
			user.Receive(message)
		}
	}
}
