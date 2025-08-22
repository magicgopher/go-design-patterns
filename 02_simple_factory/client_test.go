package simplefactory

import (
	"fmt"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestOrderLatte 模拟点单拿铁咖啡
func TestOrderLatte(t *testing.T) {
	// 创建简单工厂实例
	factory := &SimpleFactory{}
	// 创建咖啡店，传入工厂实例
	store := NewCoffeeStore(factory)

	// 通过咖啡店点单拿铁咖啡
	latte := store.OrderCoffee("latte")
	if latte != nil {
		// 打印点单成功的咖啡名称
		fmt.Printf("点了一杯 %s\n", latte.Name())
		// 为拿铁咖啡加奶
		latte.AddMilk()
		// 为拿铁咖啡加糖
		latte.AddSugar()
	} else {
		// 打印点单失败的信息
		fmt.Println("无法点单：拿铁咖啡")
	}
}

// TestOrderAmericano 模拟点单美式咖啡
func TestOrderAmericano(t *testing.T) {
	// 创建简单工厂实例
	factory := &SimpleFactory{}
	// 创建咖啡店，传入工厂实例
	store := NewCoffeeStore(factory)

	// 通过咖啡店点单美式咖啡
	americano := store.OrderCoffee("Americano")
	if americano != nil {
		// 打印点单成功的咖啡名称
		fmt.Printf("点了一杯 %s\n", americano.Name())
		// 为美式咖啡加奶
		americano.AddMilk()
		// 为美式咖啡加糖
		americano.AddSugar()
	} else {
		// 打印点单失败的信息
		fmt.Println("无法点单：美式咖啡")
	}
}

// TestOrderUnsupported 模拟点单不支持的咖啡类型
func TestOrderUnsupported(t *testing.T) {
	// 创建简单工厂实例
	factory := &SimpleFactory{}
	// 创建咖啡店，传入工厂实例
	store := NewCoffeeStore(factory)

	// 尝试点单不支持的咖啡类型（espresso）
	unknown := store.OrderCoffee("espresso")
	if unknown != nil {
		// 打印点单成功的咖啡名称（理论上不会执行）
		fmt.Printf("点了一杯 %s\n", unknown.Name())
	} else {
		// 打印点单失败的信息，提示不支持的类型
		fmt.Println("无法点单：不支持的咖啡类型 espresso")
	}
}
