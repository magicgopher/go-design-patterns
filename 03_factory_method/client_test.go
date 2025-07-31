package factorymethod

import "testing"

// 单元测试
// 模拟客户端调用

// TestOrderCoffee 测试咖啡店订购咖啡的功能。
func TestOrderCoffee(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name         string
		factory      CoffeeFactory
		expectedName string
	}{
		{
			name:         "TestLatte",
			factory:      LatteFactory{},
			expectedName: "拿铁咖啡",
		},
		{
			name:         "TestMocha",
			factory:      MochaFactory{},
			expectedName: "摩卡咖啡",
		},
		{
			name:         "TestAmericano",
			factory:      AmericanoFactory{},
			expectedName: "美式咖啡",
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建咖啡店，注入特定工厂
			coffeeStore := NewCoffeeStore(tt.factory)

			// 订购咖啡
			coffee, err := coffeeStore.OrderCoffee()
			if err != nil {
				t.Fatalf("未期望错误，但得到了错误: %v", err)
			}
			if coffee == nil {
				t.Fatalf("期望得到咖啡，但得到 nil")
			}
			if coffee.Name() != tt.expectedName {
				t.Errorf("期望的咖啡是 '%s', 但得到的是 '%s'", tt.expectedName, coffee.Name())
			}

			// 日志输出，供调试
			t.Logf("成功订购: %s", coffee.Name())
		})
	}
}

// TestCoffeeType 测试返回的咖啡对象类型。
func TestCoffeeType(t *testing.T) {
	// 定义测试用例，包含每种咖啡工厂及其期望的返回类型。
	tests := []struct {
		name         string
		factory      CoffeeFactory
		expectedType any
	}{
		{name: "TestLatteType", factory: LatteFactory{}, expectedType: &latte{}},
		{name: "TestMochaType", factory: MochaFactory{}, expectedType: &mocha{}},
		{name: "TestAmericanoType", factory: AmericanoFactory{}, expectedType: &americano{}},
	}

	// 遍历测试用例，验证每种工厂返回的咖啡对象是否为预期类型。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coffeeStore := NewCoffeeStore(tt.factory)
			coffee, err := coffeeStore.OrderCoffee()
			if err != nil {
				t.Fatalf("未期望错误，但得到了错误: %v", err)
			}
			if coffee == nil {
				t.Fatalf("期望得到咖啡，但得到 nil")
			}
			switch tt.expectedType.(type) {
			case *latte:
				if _, ok := coffee.(*latte); !ok {
					t.Errorf("期望返回 *latte 类型，但得到的是 %T", coffee)
				}
			case *mocha:
				if _, ok := coffee.(*mocha); !ok {
					t.Errorf("期望返回 *mocha 类型，但得到的是 %T", coffee)
				}
			case *americano:
				if _, ok := coffee.(*americano); !ok {
					t.Errorf("期望返回 *americano 类型，但得到的是 %T", coffee)
				}
			}
		})
	}
}
