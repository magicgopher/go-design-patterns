package bridge

import "fmt"

// 设备接口（Implementor），定义设备的操作方法
type Device interface {
	打开()
	关闭()
}

// 电视设备（Concrete Implementor），具体实现设备
type Television struct{}

// 打开 实现电视的打开操作
func (t *Television) 打开() {
	fmt.Println("电视已打开")
}

// 关闭 实现电视的关闭操作
func (t *Television) 关闭() {
	fmt.Println("电视已关闭")
}

// 空调设备（Concrete Implementor），具体实现设备
type AirConditioner struct{}

// 打开 实现空调的打开操作
func (a *AirConditioner) 打开() {
	fmt.Println("空调已打开")
}

// 关闭 实现空调的关闭操作
func (a *AirConditioner) 关闭() {
	fmt.Println("空调已关闭")
}
