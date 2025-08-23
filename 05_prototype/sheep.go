package prototype

// 原型模式

// Prototype 接口定义了克隆方法（原型接口）
type Prototype interface {
	Clone() Prototype
}

// Sheep 羊结构体（具体原型）
type Sheep struct {
	Name  string // 名字
	Age   int    // 年龄
	Color string // 颜色
}

// Clone 方法实现 Prototype 接口的克隆逻辑
func (s *Sheep) Clone() Prototype {
	return &Sheep{
		Name:  s.Name,
		Age:   s.Age,
		Color: s.Color,
	}
}
