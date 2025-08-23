package factorymethod

// 工厂方法模式

// 抽象工厂

// CoffeeFactory 定义咖啡工厂的行为
type CoffeeFactory interface {
	// CreateCoffee 创建特定类型的咖啡
	CreateCoffee() Coffee
}

// 具体工厂

// LatteFactory 生产拿铁咖啡的工厂
type LatteFactory struct{}

// NewLatteFactory 创建LatteFactory拿铁咖啡工厂实例
func NewLatteFactory() *LatteFactory {
	return &LatteFactory{}
}

// CreateCoffee 创建一个拿铁咖啡
func (lf LatteFactory) CreateCoffee() Coffee {
	return NewLatte()
}

// AmericanoFactory 生产美式咖啡的工厂
type AmericanoFactory struct{}

// CreateCoffee 创建一个美式咖啡
func (af AmericanoFactory) CreateCoffee() Coffee {
	return NewAmericano()
}

// CoffeeStore 咖啡店
type CoffeeStore struct {
	factory CoffeeFactory
}

// NewCoffeeStore 创建一个新的咖啡店
func NewCoffeeStore(factory CoffeeFactory) *CoffeeStore {
	return &CoffeeStore{factory: factory}
}

// OrderCoffee 顾客点单，返回特定工厂生产的咖啡
func (cs *CoffeeStore) OrderCoffee() Coffee {
	if cs.factory == nil {
		return nil // 防止空工厂导致 panic
	}
	return cs.factory.CreateCoffee()
}
