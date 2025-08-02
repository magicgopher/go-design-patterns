package builder

// 建造者模式

// 定义汽车相关零件

// Engine 引擎
type Engine struct {
	Type  string // 引擎类型
	Power int    // 马力
}

// Skeleton 车身骨架
type Skeleton struct {
	Material string // 材料
	Weight   int    // 重量
}

// Wheel 车轮
type Wheel struct {
	Type string // 车轮类型
	Size int    // 直径大小
}

// Car 车 表示我们需要通过建造者 builder 建造的对象
type Car struct {
	Make     string   // 制作商
	Model    string   // 型号
	Year     int      // 年份
	Color    string   // 颜色
	Engine   Engine   // 引擎
	Skeleton Skeleton // 车骨架
	Wheels   []Wheel  // 车轮
}
