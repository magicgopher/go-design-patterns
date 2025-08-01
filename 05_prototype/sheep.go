package prototype

// 原型模式

// Cloneable 定义可克隆对象的行为。
type Cloneable interface {
	// Clone 创建并返回当前对象的一个副本。
	Clone() Cloneable
}

// Sheep 表示一只羊，具有名称、年龄和类别。
type Sheep struct {
	Name     string
	Age      int
	Category string
}

// Clone 创建当前羊的副本。
func (s *Sheep) Clone() Cloneable {
	return &Sheep{
		Name:     s.Name,
		Age:      s.Age,
		Category: s.Category,
	}
}

// GetName 返回羊的名称。
func (s *Sheep) GetName() string {
	return s.Name
}

// SetName 设置羊的名称。
func (s *Sheep) SetName(name string) {
	s.Name = name
}

// GetAge 返回羊的年龄。
func (s *Sheep) GetAge() int {
	return s.Age
}

// SetAge 设置羊的年龄。
func (s *Sheep) SetAge(age int) {
	s.Age = age
}

// GetCategory 返回羊的类别。
func (s *Sheep) GetCategory() string {
	return s.Category
}

// SetCategory 设置羊的类别。
func (s *Sheep) SetCategory(category string) {
	s.Category = category
}
