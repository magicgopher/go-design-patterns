package builder

// 建造者模式

// Builder 定义抽象建造者接口，声明构建汽车各部分的抽象方法
type Builder interface {
	BuildEngine() Builder
	BuildFrame() Builder
	BuildWheels() Builder
	BuildSeats() Builder
	GetCar() Car
}

// SedanBuilder 具体建造者，用于构建轿车
type SedanBuilder struct {
	car Car
}

// NewSedanBuilder 创建一个新的轿车建造者实例
func NewSedanBuilder() *SedanBuilder {
	return &SedanBuilder{car: Car{}}
}

// BuildEngine 构建轿车的引擎，设置为汽油引擎，马力150
func (b *SedanBuilder) BuildEngine() Builder {
	b.car.Engine = Engine{
		Type:       "Gasoline",
		Horsepower: 150,
	}
	return b
}

// BuildFrame 构建轿车的车骨架，材料为钢，重量1200千克
func (b *SedanBuilder) BuildFrame() Builder {
	b.car.Frame = CarFrame{
		Material: "Steel",
		Weight:   1200,
	}
	return b
}

// BuildWheels 构建轿车的四个车轮，尺寸16英寸，材料为铝合金
func (b *SedanBuilder) BuildWheels() Builder {
	b.car.Wheels = []Wheel{
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
	}
	return b
}

// BuildSeats 构建轿车的座位，包括前排2个和后排3个，材料为皮革
func (b *SedanBuilder) BuildSeats() Builder {
	b.car.Seats = []CarSeat{
		{Count: 2, Material: "Leather"},
		{Count: 3, Material: "Leather"},
	}
	return b
}

// GetCar 返回最终构建的轿车对象
func (b *SedanBuilder) GetCar() Car {
	return b.car
}

// SUVBuilder 具体建造者，用于构建SUV
type SUVBuilder struct {
	car Car
}

// NewSUVBuilder 创建一个新的SUV建造者实例
func NewSUVBuilder() *SUVBuilder {
	return &SUVBuilder{car: Car{}}
}

// BuildEngine 构建SUV的引擎，设置为柴油引擎，马力200
func (b *SUVBuilder) BuildEngine() Builder {
	b.car.Engine = Engine{
		Type:       "Diesel",
		Horsepower: 200,
	}
	return b
}

// BuildFrame 构建SUV的车骨架，材料为铝合金，重量1800千克
func (b *SUVBuilder) BuildFrame() Builder {
	b.car.Frame = CarFrame{
		Material: "Aluminum",
		Weight:   1800,
	}
	return b
}

// BuildWheels 构建SUV的四个车轮，尺寸18英寸，材料为铝合金
func (b *SUVBuilder) BuildWheels() Builder {
	b.car.Wheels = []Wheel{
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
	}
	return b
}

// BuildSeats 构建SUV的座位，包括前排2个、后排3个、第三排2个，材料为织物
func (b *SUVBuilder) BuildSeats() Builder {
	b.car.Seats = []CarSeat{
		{Count: 2, Material: "Fabric"},
		{Count: 3, Material: "Fabric"},
		{Count: 2, Material: "Fabric"},
	}
	return b
}

// GetCar 返回最终构建的SUV对象
func (b *SUVBuilder) GetCar() Car {
	return b.car
}

// Director 指挥者，负责协调建造过程
type Director struct {
	builder Builder
}

// NewDirector 创建一个新的指挥者实例，绑定指定的建造者
func NewDirector(builderType string) *Director {
	var builder Builder
	switch builderType {
	case "sedan":
		builder = NewSedanBuilder()
	case "suv":
		builder = NewSUVBuilder()
	default:
		builder = NewSedanBuilder() // 默认使用轿车建造者
	}
	return &Director{builder: builder}
}

// Construct 按照固定顺序调用建造者的方法，构建完整的汽车
func (d *Director) Construct() Car {
	return d.builder.BuildEngine().BuildFrame().BuildWheels().BuildSeats().GetCar()
}
