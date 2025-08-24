package decorator

// 装饰者模式

// Garnish 是装饰者模式的抽象装饰者（Decorator），
// 实现 FastFood 接口并持有一个 FastFood 对象的引用，
// 为具体装饰者提供通用逻辑。
type Garnish struct {
	fastFood FastFood // 被装饰的 FastFood 对象
	price    float32  // 装饰者自身的价格
	desc     string   // 装饰者自身的描述
}

// NewGarnish 是 Garnish 的构造函数，
// 初始化被装饰的 FastFood 对象、价格和描述。
func NewGarnish(fastFood FastFood, price float32, desc string) *Garnish {
	return &Garnish{
		fastFood: fastFood,
		price:    price,
		desc:     desc,
	}
}

// GetFastFood 获取被装饰的 FastFood 对象。
func (g *Garnish) GetFastFood() FastFood {
	return g.fastFood
}

// SetFastFood 设置被装饰的 FastFood 对象。
func (g *Garnish) SetFastFood(fastFood FastFood) {
	g.fastFood = fastFood
}

// GetPrice 获取装饰者的价格。
func (g *Garnish) GetPrice() float32 {
	return g.price
}

// SetPrice 设置装饰者的价格。
func (g *Garnish) SetPrice(price float32) {
	g.price = price
}

// GetDesc 获取装饰者的描述。
func (g *Garnish) GetDesc() string {
	return g.desc
}

// SetDesc 设置装饰者的描述。
func (g *Garnish) SetDesc(desc string) {
	g.desc = desc
}

// Cost 提供默认的 Cost 实现，
// 计算装饰者自身价格与被装饰对象价格之和，
// 具体装饰者可以覆盖此方法。
func (g *Garnish) Cost() float32 {
	return g.GetPrice() + g.GetFastFood().Cost()
}

// Egg 是装饰者模式的具体装饰者（Concrete Decorator），
// 为快餐添加鸡蛋配料，扩展价格和描述。
type Egg struct {
	*Garnish // 嵌入 Garnish 以复用其功能
}

// NewEgg 是 Egg 的构造函数，
// 初始化鸡蛋装饰者，价格为 1 元，描述为“鸡蛋”。
func NewEgg(fastFood FastFood) *Egg {
	return &Egg{
		Garnish: NewGarnish(fastFood, 1.0, "鸡蛋"),
	}
}

// Cost 返回鸡蛋装饰后的总价格，
// 即鸡蛋价格加上被装饰对象的价格。
func (e *Egg) Cost() float32 {
	return e.GetPrice() + e.GetFastFood().Cost()
}

// GetDesc 返回鸡蛋装饰后的描述，
// 格式为“原描述+加鸡蛋”。
func (e *Egg) GetDesc() string {
	return e.Garnish.GetFastFood().GetDesc() + "加" + e.Garnish.GetDesc()
}

// Bacon 是装饰者模式的具体装饰者（Concrete Decorator），
// 为快餐添加培根配料，扩展价格和描述。
type Bacon struct {
	*Garnish // 嵌入 Garnish 以复用其功能
}

// NewBacon 是 Bacon 的构造函数，
// 初始化培根装饰者，价格为 2 元，描述为“培根”。
func NewBacon(fastFood FastFood) *Bacon {
	return &Bacon{
		Garnish: NewGarnish(fastFood, 2.0, "培根"),
	}
}

// Cost 返回培根装饰后的总价格，
// 即培根价格加上被装饰对象的价格。
func (b *Bacon) Cost() float32 {
	return b.GetPrice() + b.GetFastFood().Cost()
}

// GetDesc 返回培根装饰后的描述，
// 格式为“原描述+加培根”。
func (b *Bacon) GetDesc() string {
	return b.Garnish.GetFastFood().GetDesc() + "加" + b.Garnish.GetDesc()
}
