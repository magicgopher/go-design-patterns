package memento

// 备忘录模式

// 发起人

// Editor 是发起人，负责管理内容并创建/恢复备忘录
type Editor struct {
	content string // 状态：编辑器内容
}

// SetContent 设置编辑器内容
func (e *Editor) SetContent(content string) {
	e.content = content
}

// GetContent 获取编辑器内容
func (e *Editor) GetContent() string {
	return e.content
}

// CreateMemento 创建备忘录，保存当前状态
func (e *Editor) CreateMemento() *Memento {
	return &Memento{state: e.content}
}

// Restore 从备忘录恢复状态
func (e *Editor) Restore(m *Memento) {
	e.content = m.GetState()
}
