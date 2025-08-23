package factorymethod

import (
	"fmt"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestOrderLatte 模拟点单拿铁咖啡
func TestOrderLatte(t *testing.T) {
	// 创建拿铁咖啡工厂
	factory := NewLatteFactory()
	// 创建咖啡店，传入拿铁咖啡工厂
	store := NewCoffeeStore(factory)

	// 通过咖啡店点单拿铁咖啡
	latte := store.OrderCoffee()
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
	// 创建美式咖啡工厂
	factory := &AmericanoFactory{}
	// 创建咖啡店，传入美式咖啡工厂
	store := NewCoffeeStore(factory)

	// 通过咖啡店点单美式咖啡
	americano := store.OrderCoffee()
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

// TestOrderWithNilFactory 模拟使用空工厂点单
func TestOrderWithNilFactory(t *testing.T) {
	// 创建一个空工厂（模拟不支持的场景）
	var factory CoffeeFactory = nil
	// 创建咖啡店，传入空工厂
	store := NewCoffeeStore(factory)

	// 尝试点单
	coffee := store.OrderCoffee()
	if coffee != nil {
		// 打印点单成功的咖啡名称（理论上不会执行）
		fmt.Printf("点了一杯 %s\n", coffee.Name())
	} else {
		// 打印点单失败的信息，提示不支持的场景
		fmt.Println("无法点单：工厂为空")
	}
}
