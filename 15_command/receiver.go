package command

import "fmt"

// 命令模式

// Light 是接收者，知道如何执行具体操作
type Light struct {
	Name string
}

// TurnOn 打开灯
func (l *Light) TurnOn() {
	fmt.Printf("%s 灯已打开\n", l.Name)
}

// TurnOff 关闭灯
func (l *Light) TurnOff() {
	fmt.Printf("%s 灯已关闭\n", l.Name)
}
