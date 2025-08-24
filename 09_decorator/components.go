package decorator

// 装饰者模式

// FastFood 是装饰者模式的抽象组件（Component）接口，
// 定义了快餐的基本行为，包括计算价格和获取/设置描述和价格。
type FastFood interface {
	Cost() float32          // 计算总价格
	GetPrice() float32      // 获取价格
	SetPrice(price float32) // 设置价格
	GetDesc() string        // 获取描述
	SetDesc(desc string)    // 设置描述
}

// FriedRice 是装饰者模式的具体组件（Concrete Component），
// 表示一种具体的快餐——炒饭。
type FriedRice struct {
	price float32 // 价格
	desc  string  // 描述
}

// NewFriedRice 是 FriedRice 的构造函数，
// 初始化炒饭的价格为 10 元，描述为“炒饭”。
func NewFriedRice() *FriedRice {
	return &FriedRice{
		price: 10.0,
		desc:  "炒饭",
	}
}

// Cost 返回炒饭的价格，
// 实现 FastFood 接口的 Cost 方法。
func (f *FriedRice) Cost() float32 {
	return f.price
}

// GetPrice 获取炒饭的价格。
func (f *FriedRice) GetPrice() float32 {
	return f.price
}

// SetPrice 设置炒饭的价格。
func (f *FriedRice) SetPrice(price float32) {
	f.price = price
}

// GetDesc 获取炒饭的描述。
func (f *FriedRice) GetDesc() string {
	return f.desc
}

// SetDesc 设置炒饭的描述。
func (f *FriedRice) SetDesc(desc string) {
	f.desc = desc
}

// FriedNoodles 是装饰者模式的具体组件（Concrete Component），
// 表示一种具体的快餐——炒面。
type FriedNoodles struct {
	price float32 // 价格
	desc  string  // 描述
}

// NewFriedNoodles 是 FriedNoodles 的构造函数，
// 初始化炒面的价格为 12 元，描述为“炒面”。
func NewFriedNoodles() *FriedNoodles {
	return &FriedNoodles{
		price: 12.0,
		desc:  "炒面",
	}
}

// Cost 返回炒面的价格，
// 实现 FastFood 接口的 Cost 方法。
func (f *FriedNoodles) Cost() float32 {
	return f.price
}

// GetPrice 获取炒面的价格。
func (f *FriedNoodles) GetPrice() float32 {
	return f.price
}

// SetPrice 设置炒面的价格。
func (f *FriedNoodles) SetPrice(price float32) {
	f.price = price
}

// GetDesc 获取炒面的描述。
func (f *FriedNoodles) GetDesc() string {
	return f.desc
}

// SetDesc 设置炒面的描述。
func (f *FriedNoodles) SetDesc(desc string) {
	f.desc = desc
}
