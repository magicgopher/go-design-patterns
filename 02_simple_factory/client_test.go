package simplefactory

import (
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestOrderCoffee 测试咖啡店订购咖啡的功能
func TestOrderCoffee(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name         string
		coffeeType   string
		expectedName string
		expectNil    bool
	}{
		{
			name:         "TestLatte",
			coffeeType:   "latte",
			expectedName: "拿铁咖啡",
			expectNil:    false,
		},
		{
			name:         "TestAmericano",
			coffeeType:   "americano",
			expectedName: "美式咖啡",
			expectNil:    false,
		},
		{
			name:       "TestInvalidCoffeeType",
			coffeeType: "espresso",
			expectNil:  true,
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建咖啡店，注入简单工厂
			coffeeStore := NewCoffeeStore(&SimpleFactory{})

			// 订购咖啡
			coffee := coffeeStore.OrderCoffee(tt.coffeeType)

			// 检查 nil 情况
			if tt.expectNil {
				if coffee != nil {
					t.Fatalf("期望返回 nil，但得到了咖啡: %v", coffee.Name())
				}
				return
			}

			// 检查非 nil 情况
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

// TestCoffeeType 测试返回的咖啡对象类型
func TestCoffeeType(t *testing.T) {
	coffeeStore := NewCoffeeStore(&SimpleFactory{})

	// 测试拿铁咖啡类型
	t.Run("TestLatteType", func(t *testing.T) {
		coffee := coffeeStore.OrderCoffee("latte")
		if coffee == nil {
			t.Fatalf("期望得到咖啡，但得到 nil")
		}
		if _, ok := coffee.(*Latte); !ok {
			t.Errorf("期望返回 *Latte 类型，但得到的是 %T", coffee)
		}
	})

	// 测试美式咖啡类型
	t.Run("TestAmericanoType", func(t *testing.T) {
		coffee := coffeeStore.OrderCoffee("americano")
		if coffee == nil {
			t.Fatalf("期望得到咖啡，但得到 nil")
		}
		if _, ok := coffee.(*Americano); !ok {
			t.Errorf("期望返回 *Americano 类型，但得到的是 %T", coffee)
		}
	})
}
