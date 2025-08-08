package abstractfactory

// 抽象工厂模式

// 抽象产品

// Drink 定义饮品接口，包含喝的行为
type Drink interface {
	Drink() string // 返回饮用饮品的描述
}

// 具体产品

// Coffee 咖啡，实现了 Drink 接口
type Coffee struct{}

// Drink 返回饮用咖啡的描述
func (c *Coffee) Drink() string {
	return "咖啡"
}

// SoyMilk 豆浆，实现了 Drink 接口
type SoyMilk struct{}

// Drink 返回饮用豆浆的描述
func (s *SoyMilk) Drink() string {
	return "豆浆"
}
