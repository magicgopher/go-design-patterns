package factorymethod

import "testing"

// 单元测试
// 模拟客户端调用

// TestOrderCoffee 测试咖啡店订购咖啡的功能
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

			// 订购咖啡 (修正：只接收一个返回值)
			coffee := coffeeStore.OrderCoffee()
			if coffee == nil {
				t.Fatalf("期望得到咖啡，但得到 nil")
			}

			// (修正：调用导出的 Name() 方法)
			if coffee.Name() != tt.expectedName {
				t.Errorf("期望的咖啡是 '%s', 但得到的是 '%s'", tt.expectedName, coffee.Name())
			}

			// 日志输出，供调试
			t.Logf("成功订购: %s", coffee.Name())
		})
	}
}

// TestCoffeeType 测试返回的咖啡对象类型
func TestCoffeeType(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name         string
		factory      CoffeeFactory
		expectedType any
	}{
		{name: "测试Latte类型", factory: LatteFactory{}, expectedType: &Latte{}},
		{name: "测试Americano类型", factory: AmericanoFactory{}, expectedType: &Americano{}},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coffeeStore := NewCoffeeStore(tt.factory)
			coffee := coffeeStore.OrderCoffee()
			if coffee == nil {
				t.Fatalf("期望得到咖啡，但得到 nil")
			}
			switch tt.expectedType.(type) {
			case *Latte:
				if _, ok := coffee.(*Latte); !ok {
					t.Errorf("期望返回 *Latte 类型，但得到的是 %T", coffee)
				}
			case *Americano:
				if _, ok := coffee.(*Americano); !ok {
					t.Errorf("期望返回 *Americano 类型，但得到的是 %T", coffee)
				}
			}
		})
	}
}
