package prototype

import "testing"

// 单元测试
// 模拟客户端调用

// TestSheepClone 测试羊的克隆功能。
func TestSheepClone(t *testing.T) {
	// 定义测试用例，包含原型羊和期望的属性。
	tests := []struct {
		name            string
		prototype       *Sheep
		cloneName       string
		cloneAge        int
		cloneCategory   string
		modifyCloneName string
	}{
		{
			name:            "TestMerinoSheep",
			prototype:       &Sheep{Name: "Dolly", Age: 5, Category: "Merino"},
			cloneName:       "Dolly",
			cloneAge:        5,
			cloneCategory:   "Merino",
			modifyCloneName: "DollyClone",
		},
		{
			name:            "TestDorsetSheep",
			prototype:       &Sheep{Name: "Molly", Age: 3, Category: "Dorset"},
			cloneName:       "Molly",
			cloneAge:        3,
			cloneCategory:   "Dorset",
			modifyCloneName: "MollyClone",
		},
	}

	// 遍历测试用例，验证克隆羊的属性和独立性。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建管理器并添加原型羊。
			manager := NewSheepManager()
			manager.AddPrototype(tt.prototype.Name, tt.prototype)

			// 获取克隆羊。
			clone, err := manager.GetClone(tt.prototype.Name)
			if err != nil {
				t.Fatalf("未期望错误，但得到: %v", err)
			}
			if clone == nil {
				t.Fatalf("期望得到克隆羊，但得到 nil")
			}

			// 转换为 Sheep 类型并验证属性。
			clonedSheep, ok := clone.(*Sheep)
			if !ok {
				t.Fatalf("期望克隆对象为 *Sheep 类型，但得到 %T", clone)
			}
			if clonedSheep.Name != tt.cloneName {
				t.Errorf("期望克隆羊名称为 %s，但得到 %s", tt.cloneName, clonedSheep.Name)
			}
			if clonedSheep.Age != tt.cloneAge {
				t.Errorf("期望克隆羊年龄为 %d，但得到 %d", tt.cloneAge, clonedSheep.Age)
			}
			if clonedSheep.Category != tt.cloneCategory {
				t.Errorf("期望克隆羊类别为 %s，但得到 %s", tt.cloneCategory, clonedSheep.Category)
			}

			// 修改克隆羊的名称，验证原型羊未受影响。
			clonedSheep.SetName(tt.modifyCloneName)
			if tt.prototype.Name == tt.modifyCloneName {
				t.Errorf("修改克隆羊名称后，原型羊名称不应改变，但得到 %s", tt.prototype.Name)
			}

			// 日志输出，供调试。
			t.Logf("成功克隆羊: 名称=%s, 年龄=%d, 类别=%s", clonedSheep.Name, clonedSheep.Age, clonedSheep.Category)
		})
	}
}

// TestInvalidPrototype 测试克隆不存在的原型羊。
func TestInvalidPrototype(t *testing.T) {
	// 创建管理器，不添加任何原型。
	manager := NewSheepManager()

	// 尝试克隆不存在的羊。
	_, err := manager.GetClone("NonExistent")
	if err == nil {
		t.Errorf("期望得到错误，但得到 nil")
	}
	expectedErr := "原型羊 NonExistent 不存在"
	if err.Error() != expectedErr {
		t.Errorf("期望错误信息为 '%s'，但得到 '%s'", expectedErr, err.Error())
	}
}
