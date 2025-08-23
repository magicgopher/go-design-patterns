package builder

import (
	"reflect"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestSedanBuilder 测试轿车建造者的功能
func TestSedanBuilder(t *testing.T) {
	// 创建指挥者，指定轿车类型
	director := NewDirector("sedan")
	// 构建轿车
	car := director.Construct()

	// 验证引擎
	expectedEngine := Engine{Type: "Gasoline", Horsepower: 150}
	if !reflect.DeepEqual(car.Engine, expectedEngine) {
		t.Errorf("期望引擎 %v, 实际得到 %v", expectedEngine, car.Engine)
	}

	// 验证车骨架
	expectedFrame := CarFrame{Material: "Steel", Weight: 1200}
	if !reflect.DeepEqual(car.Frame, expectedFrame) {
		t.Errorf("期望车骨架 %v, 实际得到 %v", expectedFrame, car.Frame)
	}

	// 验证车轮
	expectedWheels := []Wheel{
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
		{Size: 16, Material: "Aluminum"},
	}
	if !reflect.DeepEqual(car.Wheels, expectedWheels) {
		t.Errorf("期望车轮 %v, 实际得到 %v", expectedWheels, car.Wheels)
	}

	// 验证座位
	expectedSeats := []CarSeat{
		{Count: 2, Material: "Leather"},
		{Count: 3, Material: "Leather"},
	}
	if !reflect.DeepEqual(car.Seats, expectedSeats) {
		t.Errorf("期望座位 %v, 实际得到 %v", expectedSeats, car.Seats)
	}
}

// TestSUVBuilder 测试SUV建造者的功能
func TestSUVBuilder(t *testing.T) {
	// 创建指挥者，指定SUV类型
	director := NewDirector("suv")
	// 构建SUV
	car := director.Construct()

	// 验证引擎
	expectedEngine := Engine{Type: "Diesel", Horsepower: 200}
	if !reflect.DeepEqual(car.Engine, expectedEngine) {
		t.Errorf("期望引擎 %v, 实际得到 %v", expectedEngine, car.Engine)
	}

	// 验证车骨架
	expectedFrame := CarFrame{Material: "Aluminum", Weight: 1800}
	if !reflect.DeepEqual(car.Frame, expectedFrame) {
		t.Errorf("期望车骨架 %v, 实际得到 %v", expectedFrame, car.Frame)
	}

	// 验证车轮
	expectedWheels := []Wheel{
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
		{Size: 18, Material: "Aluminum"},
	}
	if !reflect.DeepEqual(car.Wheels, expectedWheels) {
		t.Errorf("期望车轮 %v, 实际得到 %v", expectedWheels, car.Wheels)
	}

	// 验证座位
	expectedSeats := []CarSeat{
		{Count: 2, Material: "Fabric"},
		{Count: 3, Material: "Fabric"},
		{Count: 2, Material: "Fabric"},
	}
	if !reflect.DeepEqual(car.Seats, expectedSeats) {
		t.Errorf("期望座位 %v, 实际得到 %v", expectedSeats, car.Seats)
	}
}
