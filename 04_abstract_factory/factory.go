package abstractfactory

// 抽象工厂模式

// 抽象工厂

// AbstractFactory 定义抽象工厂接口，用于创建食品和饮品
type AbstractFactory interface {
	CreateFood() Food   // 创建食品
	CreateDrink() Drink // 创建饮品
}

// 具体工厂

// ChineseBreakfastFactory 中式早餐工厂，创建 Dumpling 和 SoyMilk
type ChineseBreakfastFactory struct{}

// CreateFood 创建食品，返回 Dumpling 实例
func (cbf *ChineseBreakfastFactory) CreateFood() Food {
	// 饺子
	return &Dumpling{}
}

// CreateDrink 创建饮品，返回 SoyMilk 实例
func (cbf *ChineseBreakfastFactory) CreateDrink() Drink {
	// 豆浆
	return &SoyMilk{}
}

// WesternBreakfastFactory 西式早餐工厂，创建 Bread 和 Coffee
type WesternBreakfastFactory struct{}

// CreateFood 创建食品，返回 Bread 实例
func (wbf *WesternBreakfastFactory) CreateFood() Food {
	// 面包
	return &Bread{}
}

// CreateDrink 创建饮品，返回 Coffee 实例
func (wbf *WesternBreakfastFactory) CreateDrink() Drink {
	// 咖啡
	return &Coffee{}
}
