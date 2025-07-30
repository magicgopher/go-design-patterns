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
		expectError  bool
		expectedErr  string
	}{
		{
			name:         "TestLatte",
			coffeeType:   "latte",
			expectedName: "拿铁咖啡",
			expectError:  false,
		},
		{
			name:         "TestMocha",
			coffeeType:   "mocha",
			expectedName: "摩卡咖啡",
			expectError:  false,
		},
		{
			name:         "TestAmericano",
			coffeeType:   "americano",
			expectedName: "美式咖啡",
			expectError:  false,
		},
		{
			name:        "TestInvalidCoffeeType",
			coffeeType:  "espresso",
			expectError: true,
			expectedErr: "不支持的咖啡类型: espresso",
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建咖啡店，注入简单工厂
			coffeeStore := NewCoffeeStore(&SimpleFactory{})

			// 订购咖啡
			coffee, err := coffeeStore.OrderCoffee(tt.coffeeType)

			// 检查错误情况
			if tt.expectError {
				if err == nil {
					t.Fatalf("期望返回错误，但得到了咖啡: %v", coffee)
				}
				if err.Error() != tt.expectedErr {
					t.Errorf("期望的错误信息是 '%s', 但得到的是 '%s'", tt.expectedErr, err.Error())
				}
				return
			}

			// 检查非错误情况
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

// TestCoffeeType 测试返回的咖啡对象类型
func TestCoffeeType(t *testing.T) {
	tests := []struct {
		name         string
		coffeeType   string
		expectedType interface{}
	}{
		{name: "TestLatteType", coffeeType: "latte", expectedType: &latte{}},
		{name: "TestMochaType", coffeeType: "mocha", expectedType: &mocha{}},
		{name: "TestAmericanoType", coffeeType: "americano", expectedType: &americano{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coffeeStore := NewCoffeeStore(&SimpleFactory{})
			coffee, err := coffeeStore.OrderCoffee(tt.coffeeType)
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
