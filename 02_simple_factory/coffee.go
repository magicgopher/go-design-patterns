package simplefactory

// 简单工厂 - 工厂生产的产品

// Coffee 咖啡接口
type Coffee interface {
	// Name 获取咖啡的名称
	Name() string
}

// coffeeBase 存储咖啡的通用字段
type coffeeBase struct {
	name string // 咖啡名称
}

// latte 拿铁咖啡
type latte struct {
	coffeeBase // 匿名字段
}

// Name 获取咖啡的名称
func (l latte) Name() string {
	return l.name
}

// NewLatte 创建一个新的拿铁咖啡
func NewLatte() Coffee {
	return &latte{coffeeBase{name: "拿铁咖啡"}}
}

// mocha 摩卡咖啡
type mocha struct {
	coffeeBase // 匿名字段
}

// Name 获取咖啡的名称
func (m mocha) Name() string {
	return m.name
}

// NewMocha 创建一个新的摩卡咖啡
func NewMocha() Coffee {
	return &mocha{coffeeBase{name: "摩卡咖啡"}}
}

// americano 美式咖啡
type americano struct {
	coffeeBase // 匿名字段
}

// Name 获取咖啡的名称
func (a americano) Name() string {
	return a.name
}

func NewAmericano() Coffee {
	return &americano{coffeeBase{name: "美式咖啡"}}
}
