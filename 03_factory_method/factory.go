package factorymethod

// 工厂方法 - 具体的工厂

// 工厂方法 - 抽象工厂

// CoffeeFactory 定义咖啡工厂的行为。
type CoffeeFactory interface {
	// CreateCoffee 创建特定类型的咖啡。
	CreateCoffee() (Coffee, error)
}

// 工厂方法 - 具体工厂

// LatteFactory 生产拿铁咖啡的工厂。
type LatteFactory struct{}

// CreateCoffee 创建一个拿铁咖啡。
func (lf LatteFactory) CreateCoffee() (Coffee, error) {
	return NewLatte(), nil
}

// MochaFactory 生产摩卡咖啡的工厂。
type MochaFactory struct{}

// CreateCoffee 创建一个摩卡咖啡。
func (mf MochaFactory) CreateCoffee() (Coffee, error) {
	return NewMocha(), nil
}

// AmericanoFactory 生产美式咖啡的工厂。
type AmericanoFactory struct{}

// CreateCoffee 创建一个美式咖啡。
func (af AmericanoFactory) CreateCoffee() (Coffee, error) {
	return NewAmericano(), nil
}

// CoffeeStore 咖啡店，负责处理订单。
type CoffeeStore struct {
	factory CoffeeFactory
}

// NewCoffeeStore 创建一个新的咖啡店，注入特定的咖啡工厂。
func NewCoffeeStore(factory CoffeeFactory) *CoffeeStore {
	return &CoffeeStore{factory: factory}
}

// OrderCoffee 顾客点单，返回特定工厂生产的咖啡。
func (cs *CoffeeStore) OrderCoffee() (Coffee, error) {
	return cs.factory.CreateCoffee()
}
