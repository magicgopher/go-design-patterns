package factorymethod

import "fmt"

// 工厂方法模式

// 抽象产品

// Coffee 咖啡接口
type Coffee interface {
	// 获取咖啡名称
	Name() string
	// 加奶
	AddMilk()
	// 加糖
	AddSugar()
}

// 具体产品

// Latte 拿铁咖啡
type Latte struct {
}

// NewLatte 创建Latte指针
func NewLatte() *Latte {
	return &Latte{}
}

// Name 获取咖啡名称
func (l *Latte) Name() string {
	return "拿铁咖啡"
}

// AddMilk 加奶
func (l *Latte) AddMilk() {
	fmt.Printf("%s加奶\n", l.Name())
}

// AddSugar 加糖
func (l *Latte) AddSugar() {
	fmt.Printf("%s加糖\n", l.Name())
}

// Americano 美式咖啡
type Americano struct {
}

// NewAmericano 创建Americano指针
func NewAmericano() *Americano {
	return &Americano{}
}

// Name 获取咖啡名称
func (a *Americano) Name() string {
	return "美式咖啡"
}

// AddMilk 加奶
func (a *Americano) AddMilk() {
	fmt.Printf("给%s加奶\n", a.Name())
}

// AddSugar 加糖
func (a *Americano) AddSugar() {
	fmt.Printf("给%s加糖\n", a.Name())
}
