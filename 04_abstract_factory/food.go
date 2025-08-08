package abstractfactory

// 抽象工厂模式

// 抽象产品

// Food 定义食品接口，包含吃的行为
type Food interface {
	Eat() string // 返回食用食品的描述
}

// 具体产品

// Bread 面包，实现了 Food 接口
type Bread struct{}

// Eat 返回食用面包的描述
func (b *Bread) Eat() string {
	return "面包"
}

// Dumpling 饺子，实现了 Food 接口
type Dumpling struct{}

// Eat 返回食用饺子的描述
func (d *Dumpling) Eat() string {
	return "饺子"
}
