package prototype

// 原型模式

// PrototypeManager 结构体（原型管理器角色）
// 用于存储和管理原型对象，提供克隆功能
type PrototypeManager struct {
	prototypes map[string]Prototype
}

// NewPrototypeManager 创建一个新的原型管理器
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Prototype),
	}
}

// SetPrototype 添加一个原型到管理器
func (pm *PrototypeManager) SetPrototype(name string, prototype Prototype) {
	pm.prototypes[name] = prototype
}

// GetPrototype 获取指定名称的原型的克隆副本
func (pm *PrototypeManager) GetPrototype(name string) Prototype {
	if prototype, exists := pm.prototypes[name]; exists {
		return prototype.Clone()
	}
	return nil
}
