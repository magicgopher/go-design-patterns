package abstractfactory

import "fmt"

// 抽象工厂模式

// 这里定义抽象工厂的接口和具体工厂的实现，以及咖啡店的逻辑

// ProductFactory 定义抽象工厂的行为，负责创建咖啡和甜点产品。
type ProductFactory interface {
	// CreateCoffee 创建特定类型的咖啡。
	CreateCoffee(coffeeType string) (Coffee, error)
	// CreateDessert 创建特定类型的甜点。
	CreateDessert(dessertType string) (Dessert, error)
}

// ClassicFactory 经典系列产品的工厂。
type ClassicFactory struct{}

// CreateCoffee 创建经典系列的咖啡。
func (f ClassicFactory) CreateCoffee(coffeeType string) (Coffee, error) {
	switch coffeeType {
	case "latte":
		return NewClassicLatte(), nil
	case "mocha":
		return NewClassicMocha(), nil
	case "americano":
		return NewClassicAmericano(), nil
	default:
		return nil, fmt.Errorf("不支持的咖啡类型: %v", coffeeType)
	}
}

// CreateDessert 创建经典系列的甜点。
func (f ClassicFactory) CreateDessert(dessertType string) (Dessert, error) {
	switch dessertType {
	case "cake":
		return NewClassicCake(), nil
	case "cookie":
		return NewClassicCookie(), nil
	default:
		return nil, fmt.Errorf("不支持的甜点类型: %v", dessertType)
	}
}

// ItalianFactory 意式系列产品的工厂。
type ItalianFactory struct{}

// CreateCoffee 创建意式系列的咖啡。
func (f ItalianFactory) CreateCoffee(coffeeType string) (Coffee, error) {
	switch coffeeType {
	case "latte":
		return NewItalianLatte(), nil
	case "mocha":
		return NewItalianMocha(), nil
	case "americano":
		return NewItalianAmericano(), nil
	default:
		return nil, fmt.Errorf("不支持的咖啡类型: %v", coffeeType)
	}
}

// CreateDessert 创建意式系列的甜点。
func (f ItalianFactory) CreateDessert(dessertType string) (Dessert, error) {
	switch dessertType {
	case "cake":
		return NewItalianCake(), nil
	case "cookie":
		return NewItalianCookie(), nil
	default:
		return nil, fmt.Errorf("不支持的甜点类型: %v", dessertType)
	}
}

// CoffeeStore 咖啡店，负责处理咖啡和甜点订单。
type CoffeeStore struct {
	factory ProductFactory
}

// NewCoffeeStore 创建一个新的咖啡店，注入特定的产品工厂。
func NewCoffeeStore(factory ProductFactory) *CoffeeStore {
	return &CoffeeStore{factory: factory}
}

// OrderCoffee 订购指定类型的咖啡。
func (cs *CoffeeStore) OrderCoffee(coffeeType string) (Coffee, error) {
	return cs.factory.CreateCoffee(coffeeType)
}

// OrderDessert 订购指定类型的甜点。
func (cs *CoffeeStore) OrderDessert(dessertType string) (Dessert, error) {
	return cs.factory.CreateDessert(dessertType)
}
