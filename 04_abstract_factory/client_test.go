package abstractfactory

import "testing"

// 单元测试
// 模拟客户端调用

// TestChineseBreakfastFactory 测试中式早餐工厂的创建逻辑
func TestChineseBreakfastFactory(t *testing.T) {
	factory := &ChineseBreakfastFactory{}

	// 测试创建 Food
	food := factory.CreateFood()
	if _, ok := food.(*Dumpling); !ok {
		t.Errorf("Expected food to be *Dumpling, got %T", food)
	}
	if got := food.Eat(); got != "饺子" {
		t.Errorf("Expected food.Eat() to return '饺子', got '%s'", got)
	}

	// 测试创建 Drink
	drink := factory.CreateDrink()
	if _, ok := drink.(*SoyMilk); !ok {
		t.Errorf("Expected drink to be *SoyMilk, got %T", drink)
	}
	if got := drink.Drink(); got != "豆浆" {
		t.Errorf("Drink() = %s; want 豆浆", got)
	}
}

// TestWesternBreakfastFactory 测试西式早餐工厂的创建逻辑
func TestWesternBreakfastFactory(t *testing.T) {
	factory := &WesternBreakfastFactory{}

	// 测试创建 Food
	food := factory.CreateFood()
	if _, ok := food.(*Bread); !ok {
		t.Errorf("Expected food to be *Bread, got %T", food)
	}
	if got := food.Eat(); got != "面包" {
		t.Errorf("Expected food.Eat() to return '面包', got '%s'", got)
	}

	// 测试创建 Drink
	drink := factory.CreateDrink()
	if _, ok := drink.(*Coffee); !ok {
		t.Errorf("Expected drink to be *Coffee, got %T", drink)
	}
	if got := drink.Drink(); got != "咖啡" {
		t.Errorf("Drink() = %s; want 咖啡", got)
	}
}
