package abstractfactory

// 抽象工厂模式

// 这里定义了咖啡和甜点的产品接口、以及具体的产品实现

// Coffee 定义咖啡产品的行为。
type Coffee interface {
	// Name 返回咖啡的名称。
	Name() string
}

// coffeeBase 存储咖啡的通用字段。
type coffeeBase struct {
	name string
}

// ClassicLatte 经典拿铁咖啡。
type ClassicLatte struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (l ClassicLatte) Name() string {
	return l.name
}

// NewClassicLatte 创建一个经典拿铁咖啡。
func NewClassicLatte() Coffee {
	return &ClassicLatte{coffeeBase{name: "经典拿铁咖啡"}}
}

// ClassicMocha 经典摩卡咖啡。
type ClassicMocha struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (m ClassicMocha) Name() string {
	return m.name
}

// NewClassicMocha 创建一个经典摩卡咖啡。
func NewClassicMocha() Coffee {
	return &ClassicMocha{coffeeBase{name: "经典摩卡咖啡"}}
}

// ClassicAmericano 经典美式咖啡。
type ClassicAmericano struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (a ClassicAmericano) Name() string {
	return a.name
}

// NewClassicAmericano 创建一个经典美式咖啡。
func NewClassicAmericano() Coffee {
	return &ClassicAmericano{coffeeBase{name: "经典美式咖啡"}}
}

// ItalianLatte 意式拿铁咖啡。
type ItalianLatte struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (l ItalianLatte) Name() string {
	return l.name
}

// NewItalianLatte 创建一个意式拿铁咖啡。
func NewItalianLatte() Coffee {
	return &ItalianLatte{coffeeBase{name: "意式拿铁咖啡"}}
}

// ItalianMocha 意式摩卡咖啡。
type ItalianMocha struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (m ItalianMocha) Name() string {
	return m.name
}

// NewItalianMocha 创建一个意式摩卡咖啡。
func NewItalianMocha() Coffee {
	return &ItalianMocha{coffeeBase{name: "意式摩卡咖啡"}}
}

// ItalianAmericano 意式美式咖啡。
type ItalianAmericano struct {
	coffeeBase
}

// Name 返回咖啡的名称。
func (a ItalianAmericano) Name() string {
	return a.name
}

// NewItalianAmericano 创建一个意式美式咖啡。
func NewItalianAmericano() Coffee {
	return &ItalianAmericano{coffeeBase{name: "意式美式咖啡"}}
}

// Dessert 定义甜点产品的行为。
type Dessert interface {
	// Name 返回甜点的名称。
	Name() string
}

// dessertBase 存储甜点的通用字段。
type dessertBase struct {
	name string
}

// ClassicCake 经典蛋糕。
type ClassicCake struct {
	dessertBase
}

// Name 返回甜点的名称。
func (c ClassicCake) Name() string {
	return c.name
}

// NewClassicCake 创建一个经典蛋糕。
func NewClassicCake() Dessert {
	return &ClassicCake{dessertBase{name: "经典蛋糕"}}
}

// ClassicCookie 经典曲奇。
type ClassicCookie struct {
	dessertBase
}

// Name 返回甜点的名称。
func (c ClassicCookie) Name() string {
	return c.name
}

// NewClassicCookie 创建一个经典曲奇。
func NewClassicCookie() Dessert {
	return &ClassicCookie{dessertBase{name: "经典曲奇"}}
}

// ItalianCake 意式蛋糕。
type ItalianCake struct {
	dessertBase
}

// Name 返回甜点的名称。
func (c ItalianCake) Name() string {
	return c.name
}

// NewItalianCake 创建一个意式蛋糕。
func NewItalianCake() Dessert {
	return &ItalianCake{dessertBase{name: "意式蛋糕"}}
}

// ItalianCookie 意式曲奇。
type ItalianCookie struct {
	dessertBase
}

// Name 返回甜点的名称。
func (c ItalianCookie) Name() string {
	return c.name
}

// NewItalianCookie 创建一个意式曲奇。
func NewItalianCookie() Dessert {
	return &ItalianCookie{dessertBase{name: "意式曲奇"}}
}
