package decorator

import (
	"fmt"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestFastFood 测试装饰者模式的各种场景
func TestFastFood(t *testing.T) {
	// 定义测试用例，包含测试名称、快餐对象、预期描述和预期价格
	tests := []struct {
		name         string   // 测试用例名称
		food         FastFood // 快餐对象
		expectedDesc string   // 预期描述
		expectedCost float32  // 预期价格
	}{
		{
			name:         "FriedRice", // 测试基础炒饭
			food:         NewFriedRice(),
			expectedDesc: "炒饭",
			expectedCost: 10.0,
		},
		{
			name:         "FriedRice with Egg", // 测试加鸡蛋的炒饭
			food:         NewEgg(NewFriedRice()),
			expectedDesc: "炒饭加鸡蛋",
			expectedCost: 11.0,
		},
		{
			name:         "FriedNoodles with Bacon", // 测试加培根的炒面
			food:         NewBacon(NewFriedNoodles()),
			expectedDesc: "炒面加培根",
			expectedCost: 14.0,
		},
		{
			name:         "FriedRice with Egg and Bacon", // 测试加鸡蛋和培根的炒饭
			food:         NewBacon(NewEgg(NewFriedRice())),
			expectedDesc: "炒饭加鸡蛋加培根",
			expectedCost: 13.0,
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 获取实际的描述和价格
			desc := tt.food.GetDesc()
			cost := tt.food.Cost()

			// 验证描述是否符合预期
			if desc != tt.expectedDesc {
				t.Errorf("预期 desc %q，实际为 %q\n", tt.expectedDesc, desc)
			}
			// 验证价格是否符合预期
			if cost != tt.expectedCost {
				t.Errorf("预期成本 %.1f，实际收益 %.1f\n", tt.expectedCost, cost)
			}

			// 输出测试结果，便于调试
			t.Logf("%s %.1f元", desc, cost)
		})
	}
}

// TestMain 演示装饰者模式的使用，
// 通过不同的装饰组合展示如何为快餐添加配料。
func TestMain(t *testing.T) {
	// 场景 1：简单的炒饭
	friedRice := NewFriedRice()
	fmt.Printf("场景 1：%s，价格：%.2f\n", friedRice.GetDesc(), friedRice.Cost())

	// 场景 2：炒饭加鸡蛋
	eggFriedRice := NewEgg(friedRice)
	fmt.Printf("场景 2：%s，价格：%.2f\n", eggFriedRice.GetDesc(), eggFriedRice.Cost())

	// 场景 3：炒饭加鸡蛋和培根
	baconEggFriedRice := NewBacon(eggFriedRice)
	fmt.Printf("场景 3：%s，价格：%.2f\n", baconEggFriedRice.GetDesc(), baconEggFriedRice.Cost())

	// 场景 4：炒面
	friedNoodles := NewFriedNoodles()
	fmt.Printf("场景 4：%s，价格：%.2f\n", friedNoodles.GetDesc(), friedNoodles.Cost())

	// 场景 5：炒面加培根
	baconFriedNoodles := NewBacon(friedNoodles)
	fmt.Printf("场景 5：%s，价格：%.2f\n", baconFriedNoodles.GetDesc(), baconFriedNoodles.Cost())

	// 场景 6：炒面加双份鸡蛋
	eggFriedNoodles := NewEgg(friedNoodles)
	doubleEggFriedNoodles := NewEgg(eggFriedNoodles)
	fmt.Printf("场景 6：%s，价格：%.2f\n", doubleEggFriedNoodles.GetDesc(), doubleEggFriedNoodles.Cost())
}
