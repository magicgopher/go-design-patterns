package builder

// 建造者模式

// 构建的复杂产品

// Engine 表示汽车的引擎结构体
type Engine struct {
	Type       string // 引擎类型（如汽油、电动）
	Horsepower int    // 马力，单位为马力（hp）
}

// CarFrame 表示汽车的车骨架结构体
type CarFrame struct {
	Material string // 骨架材料（如钢、铝合金）
	Weight   int    // 骨架重量，单位为千克（kg）
}

// Wheel 表示汽车的车轮结构体
type Wheel struct {
	Size     int    // 车轮尺寸，单位为英寸
	Material string // 车轮材料（如铝合金、钢）
}

// CarSeat 表示汽车的座位结构体
type CarSeat struct {
	Count    int    // 座位数量
	Material string // 座椅材料（如皮革、织物）
}

// Car 表示最终的汽车产品结构体，包含引擎、车骨架、车轮和座位
type Car struct {
	Engine Engine    // 汽车的引擎
	Frame  CarFrame  // 汽车的车骨架
	Wheels []Wheel   // 汽车的车轮列表
	Seats  []CarSeat // 汽车的座位列表
}
