package memento

// 备忘录模式

// 备忘录

// Memento 是备忘录，存储编辑器的状态
type Memento struct {
	state string // 状态：编辑器内容的快照
}

// GetState 获取备忘录中的状态
func (m *Memento) GetState() string {
	return m.state
}
