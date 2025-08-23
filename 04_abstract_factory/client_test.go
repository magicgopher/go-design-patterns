package abstractfactory

import "testing"

// 单元测试
// 模拟客户端调用

// TestChineseBreakfast 测试中式早餐工厂
func TestChineseBreakfast(t *testing.T) {
	// 创建中式早餐工厂
	factory := &ChineseBreakfastFactory{}

	// 使用工厂创建食品和饮品
	food := factory.CreateFood()
	drink := factory.CreateDrink()

	// 验证食品和饮品是否正确
	if food.Eat() != "饺子" {
		t.Errorf("期望食品是 %s，但实际得到 %s", "饺子", food.Eat())
	}
	if drink.Drink() != "豆浆" {
		t.Errorf("期望饮品是 %s，但实际得到 %s", "豆浆", drink.Drink())
	}
}

// TestWesternBreakfast 测试西式早餐工厂
func TestWesternBreakfast(t *testing.T) {
	// 创建西式早餐工厂
	factory := &WesternBreakfastFactory{}

	// 使用工厂创建食品和饮品
	food := factory.CreateFood()
	drink := factory.CreateDrink()

	// 验证食品和饮品是否正确
	if food.Eat() != "面包" {
		t.Errorf("期望食品是 %s，但实际得到 %s", "面包", food.Eat())
	}
	if drink.Drink() != "咖啡" {
		t.Errorf("期望饮品是 %s，但实际得到 %s", "咖啡", drink.Drink())
	}
}

// TestAbstractFactoryInterface 测试抽象工厂接口的通用性
func TestAbstractFactoryInterface(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name      string
		factory   AbstractFactory
		wantFood  string
		wantDrink string
	}{
		{
			name:      "中式早餐",
			factory:   &ChineseBreakfastFactory{},
			wantFood:  "饺子",
			wantDrink: "豆浆",
		},
		{
			name:      "西式早餐",
			factory:   &WesternBreakfastFactory{},
			wantFood:  "面包",
			wantDrink: "咖啡",
		},
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			food := tt.factory.CreateFood()
			drink := tt.factory.CreateDrink()

			if food.Eat() != tt.wantFood {
				t.Errorf("期望食品是 %s，但实际得到 %s", tt.wantFood, food.Eat())
			}
			if drink.Drink() != tt.wantDrink {
				t.Errorf("期望饮品是 %s，但实际得到 %s", tt.wantDrink, drink.Drink())
			}
		})
	}
}
