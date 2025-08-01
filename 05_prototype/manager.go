package prototype

import "fmt"

// 原型模式
// 用于管理原型羊并提供克隆功能

// SheepManager 管理原型羊并提供克隆功能。
type SheepManager struct {
	prototypes map[string]*Sheep
}

// NewSheepManager 创建一个新的羊管理器。
func NewSheepManager() *SheepManager {
	return &SheepManager{
		prototypes: make(map[string]*Sheep),
	}
}

// AddPrototype 添加一个原型羊到管理器。
func (sm *SheepManager) AddPrototype(name string, sheep *Sheep) {
	sm.prototypes[name] = sheep
}

// GetClone 根据名称获取原型羊的克隆。
func (sm *SheepManager) GetClone(name string) (Cloneable, error) {
	prototype, exists := sm.prototypes[name]
	if !exists {
		return nil, fmt.Errorf("原型羊 %s 不存在", name)
	}
	return prototype.Clone(), nil
}
