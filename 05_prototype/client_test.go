package prototype

import "testing"

// 单元测试
// 模拟客户端调用

// TestCloneSheep 测试羊的克隆功能
func TestCloneSheep(t *testing.T) {
	// 创建一个原型羊
	originalSheep := &Sheep{
		Name:  "Dolly",
		Age:   2,
		Color: "White",
	}

	// 客户端直接调用 Clone 方法（客户端角色）
	clonedSheep := originalSheep.Clone()

	// 验证克隆结果
	if clonedSheep == originalSheep {
		t.Error("克隆的羊与原羊是同一对象，克隆失败")
	}
	sheep, ok := clonedSheep.(*Sheep)
	if !ok {
		t.Error("克隆对象类型错误")
	}
	if sheep.Name != originalSheep.Name || sheep.Age != originalSheep.Age || sheep.Color != originalSheep.Color {
		t.Error("克隆的羊属性不一致")
	}
}

// TestPrototypeManager 测试原型管理器的功能
func TestPrototypeManager(t *testing.T) {
	// 创建原型管理器
	manager := NewPrototypeManager()

	// 创建并注册原型羊
	dolly := &Sheep{
		Name:  "Dolly",
		Age:   2,
		Color: "White",
	}
	manager.SetPrototype("Dolly", dolly)

	// 客户端通过管理器获取克隆羊（客户端角色）
	clonedPrototype := manager.GetPrototype("Dolly")
	if clonedPrototype == nil {
		t.Error("无法获取克隆羊")
	}

	// 验证克隆结果
	clonedSheep, ok := clonedPrototype.(*Sheep)
	if !ok {
		t.Error("克隆对象类型错误")
	}
	if clonedSheep == dolly {
		t.Error("克隆的羊与原羊是同一对象，克隆失败")
	}
	if clonedSheep.Name != dolly.Name || clonedSheep.Age != dolly.Age || clonedSheep.Color != dolly.Color {
		t.Error("克隆的羊属性不一致")
	}
}
