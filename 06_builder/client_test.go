package builder

import (
	"reflect"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestCarBuilder 测试建造者模式构建汽车的功能。
func TestCarBuilder(t *testing.T) {
	// 定义测试用例，包含输入参数和期望的汽车属性，使用中文数据。
	tests := []struct {
		name         string
		make         string
		model        string
		year         int
		color        string
		engine       Engine
		skeleton     Skeleton
		wheels       []Wheel
		expectedCar  *Car
		modifyClone  bool   // 是否修改克隆对象以测试深拷贝
		modifiedMake string // 修改后的制造商名称
	}{
		{
			name:  "TestSportsCar",
			make:  "法拉利",
			model: "488 GTB",
			year:  2023,
			color: "红色",
			engine: Engine{
				Type:  "V8引擎",
				Power: 670,
			},
			skeleton: Skeleton{
				Material: "碳纤维",
				Weight:   1200,
			},
			wheels: []Wheel{
				{Type: "合金轮毂", Size: 20},
				{Type: "合金轮毂", Size: 20},
				{Type: "合金轮毂", Size: 20},
				{Type: "合金轮毂", Size: 20},
			},
			expectedCar: &Car{
				Make:  "法拉利",
				Model: "488 GTB",
				Year:  2023,
				Color: "红色",
				Engine: Engine{
					Type:  "V8引擎",
					Power: 670,
				},
				Skeleton: Skeleton{
					Material: "碳纤维",
					Weight:   1200,
				},
				Wheels: []Wheel{
					{Type: "合金轮毂", Size: 20},
					{Type: "合金轮毂", Size: 20},
					{Type: "合金轮毂", Size: 20},
					{Type: "合金轮毂", Size: 20},
				},
			},
			modifyClone:  true,
			modifiedMake: "改装法拉利",
		},
		{
			name:  "TestSUV",
			make:  "丰田",
			model: "RAV4",
			year:  2022,
			color: "银色",
			engine: Engine{
				Type:  "四缸引擎",
				Power: 203,
			},
			skeleton: Skeleton{
				Material: "钢材",
				Weight:   1600,
			},
			wheels: []Wheel{
				{Type: "钢制轮毂", Size: 17},
				{Type: "钢制轮毂", Size: 17},
				{Type: "钢制轮毂", Size: 17},
				{Type: "钢制轮毂", Size: 17},
			},
			expectedCar: &Car{
				Make:  "丰田",
				Model: "RAV4",
				Year:  2022,
				Color: "银色",
				Engine: Engine{
					Type:  "四缸引擎",
					Power: 203,
				},
				Skeleton: Skeleton{
					Material: "钢材",
					Weight:   1600,
				},
				Wheels: []Wheel{
					{Type: "钢制轮毂", Size: 17},
					{Type: "钢制轮毂", Size: 17},
					{Type: "钢制轮毂", Size: 17},
					{Type: "钢制轮毂", Size: 17},
				},
			},
			modifyClone: false,
		},
	}

	// 遍历测试用例，验证建造者和导向器的构建结果。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建建造者。
			builder := NewConcreteCarBuilder()
			// 创建导向器。
			director := NewDirector(builder)

			// 使用导向器构建汽车。
			car := director.Construct(
				tt.make,
				tt.model,
				tt.year,
				tt.color,
				tt.engine,
				tt.skeleton,
				tt.wheels,
			)

			// 验证构建的汽车是否符合预期。
			if !reflect.DeepEqual(car, tt.expectedCar) {
				t.Errorf("期望的汽车为 %+v，但得到 %+v", tt.expectedCar, car)
			}

			// 测试深拷贝：修改克隆对象的属性，验证原建造者状态未变。
			if tt.modifyClone {
				car.Make = tt.modifiedMake
				secondCar := builder.Build()
				if secondCar.Make == tt.modifiedMake {
					t.Errorf("修改克隆对象后，建造者状态不应改变，但得到 Make=%s", secondCar.Make)
				}
				if secondCar.Make != tt.make {
					t.Errorf("期望建造者保留原始 Make=%s，但得到 %s", tt.make, secondCar.Make)
				}
			}

			// 测试重置功能：重置后构建新车应为空状态。
			builder.Reset()
			emptyCar := builder.Build()
			if emptyCar.Make != "" || emptyCar.Model != "" || len(emptyCar.Wheels) != 0 {
				t.Errorf("重置后期望空汽车，但得到 %+v", emptyCar)
			}

			// 日志输出，供调试。
			t.Logf("成功构建汽车: 制造商=%s, 型号=%s, 年份=%d, 颜色=%s", car.Make, car.Model, car.Year, car.Color)
		})
	}
}

// TestCarBuilderWithoutDirector 测试不使用导向器直接使用建造者。
func TestCarBuilderWithoutDirector(t *testing.T) {
	// 创建建造者。
	builder := NewConcreteCarBuilder()

	// 手动设置属性，使用中文数据。
	car := builder.
		SetMake("本田").
		SetModel("思域").
		SetYear(2021).
		SetColor("蓝色").
		SetEngine(Engine{Type: "四缸引擎", Power: 158}).
		SetSkeleton(Skeleton{Material: "钢材", Weight: 1300}).
		AddWheel(Wheel{Type: "合金轮毂", Size: 16}).
		AddWheel(Wheel{Type: "合金轮毂", Size: 16}).
		AddWheel(Wheel{Type: "合金轮毂", Size: 16}).
		AddWheel(Wheel{Type: "合金轮毂", Size: 16}).
		Build()

	// 定义期望的汽车。
	expectedCar := &Car{
		Make:  "本田",
		Model: "思域",
		Year:  2021,
		Color: "蓝色",
		Engine: Engine{
			Type:  "四缸引擎",
			Power: 158,
		},
		Skeleton: Skeleton{
			Material: "钢材",
			Weight:   1300,
		},
		Wheels: []Wheel{
			{Type: "合金轮毂", Size: 16},
			{Type: "合金轮毂", Size: 16},
			{Type: "合金轮毂", Size: 16},
			{Type: "合金轮毂", Size: 16},
		},
	}

	// 验证构建的汽车。
	if !reflect.DeepEqual(car, expectedCar) {
		t.Errorf("期望的汽车为 %+v，但得到 %+v", expectedCar, car)
	}

	// 日志输出，供调试。
	t.Logf("成功构建汽车: 制造商=%s, 型号=%s, 年份=%d, 颜色=%s", car.Make, car.Model, car.Year, car.Color)
}

// TestCarBuilderInvalidInput 测试无效输入的情况。
func TestCarBuilderInvalidInput(t *testing.T) {
	// 创建建造者和导向器。
	builder := NewConcreteCarBuilder()
	director := NewDirector(builder)

	// 测试无效年份（负数）。
	car := director.Construct(
		"宝马",
		"X5",
		-1, // 无效年份
		"黑色",
		Engine{Type: "V6引擎", Power: 335},
		Skeleton{Material: "铝合金", Weight: 1800},
		[]Wheel{
			{Type: "合金轮毂", Size: 19},
			{Type: "合金轮毂", Size: 19},
			{Type: "合金轮毂", Size: 19},
			{Type: "合金轮毂", Size: 19},
		},
	)

	// 验证年份是否被设置（当前实现不检查有效性，记录潜在改进）。
	if car.Year < 0 {
		t.Logf("警告：允许设置无效年份 %d，建议在建造者中添加验证", car.Year)
	}

	// 测试少于 4 个车轮。
	builder.Reset()
	car = builder.
		SetMake("特斯拉").
		SetModel("Model 3").
		SetYear(2020).
		SetColor("白色").
		SetEngine(Engine{Type: "电动引擎", Power: 225}).
		SetSkeleton(Skeleton{Material: "钢材", Weight: 1400}).
		AddWheel(Wheel{Type: "空气动力轮毂", Size: 18}).
		AddWheel(Wheel{Type: "空气动力轮毂", Size: 18}).
		Build()

	// 验证车轮数量（当前实现不检查，记录潜在改进）。
	if len(car.Wheels) < 4 {
		t.Logf("警告：车轮数量 %d 小于 4，建议在 Build 中添加验证", len(car.Wheels))
	}
}
