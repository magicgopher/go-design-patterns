package facade

import "fmt"

// 外观模式

// 子系统角色

// Projector 是子系统组件，负责投影仪操作
type Projector struct{}

func (p *Projector) TurnOn() {
	fmt.Println("投影仪已开启")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("投影仪输入设置为: %s\n", input)
}

func (p *Projector) TurnOff() {
	fmt.Println("投影仪已关闭")
}
