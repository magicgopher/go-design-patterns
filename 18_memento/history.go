package memento

// 备忘录模式

// 管理者

// History 是管理者，负责存储和管理备忘录
type History struct {
	mementos []*Memento
}

// AddMemento 添加备忘录到历史记录
func (h *History) AddMemento(m *Memento) {
	h.mementos = append(h.mementos, m)
}

// GetMemento 获取指定索引的备忘录
func (h *History) GetMemento(index int) *Memento {
	if index < 0 || index >= len(h.mementos) {
		return nil
	}
	return h.mementos[index]
}
