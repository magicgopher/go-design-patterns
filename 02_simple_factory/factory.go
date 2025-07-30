package simplefactory

import "fmt"

// 简单工厂 - 工厂

// CoffeeFactory 定义咖啡工厂的行为。
type CoffeeFactory interface {
	CreateCoffee(coffeeType string) (Coffee, error)
}

// SimpleFactory 是一个简单的咖啡工厂。
type SimpleFactory struct{}

// CreateCoffee 根据类型创建咖啡。
func (sf SimpleFactory) CreateCoffee(coffeeType string) (Coffee, error) {
	switch coffeeType {
	case "latte":
		return NewLatte(), nil
	case "americano":
		return NewAmericano(), nil
	case "mocha":
		return NewMocha(), nil
	default:
		return nil, fmt.Errorf("不支持的咖啡类型: %v", coffeeType)
	}
}

// CoffeeStore 咖啡店，负责处理订单。
type CoffeeStore struct {
	factory CoffeeFactory
}

// NewCoffeeStore 创建一个新的咖啡店。
func NewCoffeeStore(factory CoffeeFactory) *CoffeeStore {
	return &CoffeeStore{factory: factory}
}

// OrderCoffee 顾客点单。
func (cs *CoffeeStore) OrderCoffee(coffeeType string) (Coffee, error) {
	return cs.factory.CreateCoffee(coffeeType)
}
