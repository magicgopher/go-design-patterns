package simplefactory

// 简单工厂 - 工厂

// CoffeeFactory 定义咖啡工厂的行为
type CoffeeFactory interface {
	// 制作咖啡
	// coffeeType 咖啡类型
	CreateCoffee(coffeeType string) Coffee
}

// SimpleFactory 是一个简单的咖啡工厂
type SimpleFactory struct{}

// CreateCoffee 根据指定的类型创建咖啡
func (sf *SimpleFactory) CreateCoffee(coffeeType string) Coffee {
	switch coffeeType {
	case "latte", "Latte":
		return NewLatte()
	case "americano", "Americano":
		return NewAmericano()
	default:
		return nil // 不支持的类型返回 nil
	}
}

// CoffeeStore 咖啡店
type CoffeeStore struct {
	factory *SimpleFactory
}

// NewCoffeeStore 创建一个新的咖啡店
func NewCoffeeStore(factory *SimpleFactory) *CoffeeStore {
	return &CoffeeStore{
		factory: factory,
	}
}

// OrderCoffee 顾客点单
func (cs *CoffeeStore) OrderCoffee(coffeeType string) Coffee {
	coffee := cs.factory.CreateCoffee(coffeeType)
	// 可以在这里添加一些通用的准备步骤
	return coffee
}
