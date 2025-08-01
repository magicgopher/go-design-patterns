package abstractfactory

import "testing"

// 单元测试
// 模拟客户端调用

// TestOrderProducts 测试咖啡店订购咖啡和甜点的功能。
func TestOrderProducts(t *testing.T) {
	// 定义测试用例，包含工厂、产品类型及期望的名称。
	tests := []struct {
		name               string
		factory            ProductFactory
		coffeeType         string
		dessertType        string
		expectedCoffee     string
		expectedDessert    string
		expectCoffeeErr    bool
		expectDessertErr   bool
		expectedCoffeeErr  string
		expectedDessertErr string
	}{
		{
			name:             "TestClassicLatteAndCake",
			factory:          ClassicFactory{},
			coffeeType:       "latte",
			dessertType:      "cake",
			expectedCoffee:   "经典拿铁咖啡",
			expectedDessert:  "经典蛋糕",
			expectCoffeeErr:  false,
			expectDessertErr: false,
		},
		{
			name:             "TestClassicMochaAndCookie",
			factory:          ClassicFactory{},
			coffeeType:       "mocha",
			dessertType:      "cookie",
			expectedCoffee:   "经典摩卡咖啡",
			expectedDessert:  "经典曲奇",
			expectCoffeeErr:  false,
			expectDessertErr: false,
		},
		{
			name:             "TestItalianAmericanoAndCake",
			factory:          ItalianFactory{},
			coffeeType:       "americano",
			dessertType:      "cake",
			expectedCoffee:   "意式美式咖啡",
			expectedDessert:  "意式蛋糕",
			expectCoffeeErr:  false,
			expectDessertErr: false,
		},
		{
			name:              "TestInvalidCoffeeType",
			factory:           ClassicFactory{},
			coffeeType:        "espresso",
			dessertType:       "cake",
			expectedDessert:   "经典蛋糕",
			expectCoffeeErr:   true,
			expectDessertErr:  false,
			expectedCoffeeErr: "不支持的咖啡类型: espresso",
		},
		{
			name:               "TestInvalidDessertType",
			factory:            ItalianFactory{},
			coffeeType:         "latte",
			dessertType:        "pie",
			expectedCoffee:     "意式拿铁咖啡",
			expectCoffeeErr:    false,
			expectDessertErr:   true,
			expectedDessertErr: "不支持的甜点类型: pie",
		},
	}

	// 遍历测试用例，验证咖啡和甜点的订购结果。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建咖啡店，注入特定工厂。
			store := NewCoffeeStore(tt.factory)

			// 订购咖啡。
			coffee, coffeeErr := store.OrderCoffee(tt.coffeeType)
			if tt.expectCoffeeErr {
				if coffeeErr == nil {
					t.Errorf("期望咖啡订购错误，但得到 nil")
				} else if coffeeErr.Error() != tt.expectedCoffeeErr {
					t.Errorf("期望的咖啡错误是 '%s', 但得到 '%s'", tt.expectedCoffeeErr, coffeeErr.Error())
				}
			} else {
				if coffeeErr != nil {
					t.Errorf("未期望咖啡错误，但得到: %v", coffeeErr)
				} else if coffee == nil {
					t.Errorf("期望得到咖啡，但得到 nil")
				} else if coffee.Name() != tt.expectedCoffee {
					t.Errorf("期望的咖啡名称是 '%s', 但得到 '%s'", tt.expectedCoffee, coffee.Name())
				}
			}

			// 订购甜点。
			dessert, dessertErr := store.OrderDessert(tt.dessertType)
			if tt.expectDessertErr {
				if dessertErr == nil {
					t.Errorf("期望甜点订购错误，但得到 nil")
				} else if dessertErr.Error() != tt.expectedDessertErr {
					t.Errorf("期望的甜点错误是 '%s', 但得到 '%s'", tt.expectedDessertErr, dessertErr.Error())
				}
			} else {
				if dessertErr != nil {
					t.Errorf("未期望甜点错误，但得到: %v", dessertErr)
				} else if dessert == nil {
					t.Errorf("期望得到甜点，但得到 nil")
				} else if dessert.Name() != tt.expectedDessert {
					t.Errorf("期望的甜点名称是 '%s', 但得到 '%s'", tt.expectedDessert, dessert.Name())
				}
			}

			// 日志输出，供调试。
			if coffee != nil && dessert != nil {
				t.Logf("成功订购: 咖啡 '%s', 甜点 '%s'", coffee.Name(), dessert.Name())
			}
		})
	}
}

// TestProductTypes 测试返回的咖啡和甜点对象的类型。
func TestProductTypes(t *testing.T) {
	// 定义测试用例，包含工厂及期望的类型。
	tests := []struct {
		name                string
		factory             ProductFactory
		coffeeType          string
		dessertType         string
		expectedCoffeeType  any
		expectedDessertType any
	}{
		{
			name:                "TestClassicLatteAndCake",
			factory:             ClassicFactory{},
			coffeeType:          "latte",
			dessertType:         "cake",
			expectedCoffeeType:  &ClassicLatte{},
			expectedDessertType: &ClassicCake{},
		},
		{
			name:                "TestItalianMochaAndCookie",
			factory:             ItalianFactory{},
			coffeeType:          "mocha",
			dessertType:         "cookie",
			expectedCoffeeType:  &ItalianMocha{},
			expectedDessertType: &ItalianCookie{},
		},
	}

	// 遍历测试用案，验证返回的咖啡和甜点类型。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建咖啡店，注入特定工厂。
			store := NewCoffeeStore(tt.factory)

			// 订购咖啡并验证类型。
			coffee, err := store.OrderCoffee(tt.coffeeType)
			if err != nil {
				t.Fatalf("未期望咖啡错误，但得到: %v", err)
			}
			if coffee == nil {
				t.Fatalf("期望得到咖啡，但得到 nil")
			}
			switch tt.expectedCoffeeType.(type) {
			case *ClassicLatte:
				if _, ok := coffee.(*ClassicLatte); !ok {
					t.Errorf("期望返回 *ClassicLatte 类型，但得到 %T", coffee)
				}
			case *ItalianMocha:
				if _, ok := coffee.(*ItalianMocha); !ok {
					t.Errorf("期望返回 *ItalianMocha 类型，但得到 %T", coffee)
				}
			}

			// 订购甜点并验证类型。
			dessert, err := store.OrderDessert(tt.dessertType)
			if err != nil {
				t.Fatalf("未期望甜点错误，但得到: %v", err)
			}
			if dessert == nil {
				t.Fatalf("期望得到甜点，但得到 nil")
			}
			switch tt.expectedDessertType.(type) {
			case *ClassicCake:
				if _, ok := dessert.(*ClassicCake); !ok {
					t.Errorf("期望返回 *ClassicCake 类型，但得到 %T", dessert)
				}
			case *ItalianCookie:
				if _, ok := dessert.(*ItalianCookie); !ok {
					t.Errorf("期望返回 *ItalianCookie 类型，但得到 %T", dessert)
				}
			}
		})
	}
}
