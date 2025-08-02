package builder

// 建造者模式

// 定义建造者 builder

// CarBuilder 是构建汽车的接口
type CarBuilder interface {
	// SetMake 设置汽车的制造商
	SetMake(string) CarBuilder
	// SetModel 设置汽车的型号
	SetModel(string) CarBuilder
	// SetYear 设置汽车的年份
	SetYear(int) CarBuilder
	// SetColor 设置汽车的颜色
	SetColor(string) CarBuilder
	// SetEngine 设置汽车的引擎
	SetEngine(Engine) CarBuilder
	// SetSkeleton 设置汽车的车身骨架
	SetSkeleton(Skeleton) CarBuilder
	// AddWheel 添加车轮
	AddWheel(Wheel) CarBuilder
	// Build 返回构建好的汽车
	Build() *Car
}

// ConcreteCarBuilder 是具体的汽车构建器实现
type ConcreteCarBuilder struct {
	car *Car // 汽车结构体指针
}

// NewConcreteCarBuilder 创建一个新的具体汽车构建器
func NewConcreteCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{car: &Car{Wheels: []Wheel{}}}
}

// SetMake 设置汽车的制造商
func (b *ConcreteCarBuilder) SetMake(make string) CarBuilder {
	b.car.Make = make
	return b
}

// SetModel 设置汽车的型号
func (b *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	b.car.Model = model
	return b
}

// SetYear 设置汽车的年份
func (b *ConcreteCarBuilder) SetYear(year int) CarBuilder {
	b.car.Year = year
	return b
}

// SetColor 设置汽车的颜色
func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.car.Color = color
	return b
}

// SetEngine 设置汽车的引擎
func (b *ConcreteCarBuilder) SetEngine(engine Engine) CarBuilder {
	b.car.Engine = engine
	return b
}

// SetSkeleton 设置汽车的车身骨架
func (b *ConcreteCarBuilder) SetSkeleton(skeleton Skeleton) CarBuilder {
	b.car.Skeleton = skeleton
	return b
}

// AddWheel 添加车轮
func (b *ConcreteCarBuilder) AddWheel(wheel Wheel) CarBuilder {
	b.car.Wheels = append(b.car.Wheels, wheel)
	return b
}

// Build 返回构建好的汽车
//func (b *ConcreteCarBuilder) Build() *Car {
//	return b.car
//}

// Build 返回构建好的汽车
func (b *ConcreteCarBuilder) Build() *Car {
	car := *b.car // 复制 Car
	car.Wheels = make([]Wheel, len(b.car.Wheels))
	copy(car.Wheels, b.car.Wheels) // 复制 Wheels 切片
	return &car
}

// Director 用于管理汽车的构建过程
type Director struct {
	builder CarBuilder
}

// NewDirector 创建一个新的导向器
func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

// Construct 指导构建汽车的过程
// 该方法接收汽车的各个组成部分作为参数，并返回构建好的汽车对象
func (d *Director) Construct(make, model string, year int, color string, engine Engine, skeleton Skeleton, wheels []Wheel) *Car {
	// 设置汽车的基本属性
	d.builder.SetMake(make).
		SetModel(model).
		SetYear(year).
		SetColor(color).
		SetEngine(engine).
		SetSkeleton(skeleton)

	// 设置车轮
	for _, wheel := range wheels {
		d.builder.AddWheel(wheel)
	}

	// 返回构建好的汽车
	return d.builder.Build()
}

// Reset 重置建造者
func (b *ConcreteCarBuilder) Reset() {
	b.car = &Car{Wheels: []Wheel{}}
}
