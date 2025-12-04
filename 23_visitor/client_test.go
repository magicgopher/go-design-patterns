package visitor

import (
	"math"
	"testing"
)

// ObjectStructure 对象结构
// 在实际应用中，这可能是一个复杂的组合结构或列表
type ShapeCollection struct {
	shapes []Shape
}

// Add 添加元素
func (sc *ShapeCollection) Add(s Shape) {
	sc.shapes = append(sc.shapes, s)
}

// Accept 遍历元素并调用它们的 Accept 方法
func (sc *ShapeCollection) Accept(v Visitor) {
	for _, shape := range sc.shapes {
		shape.Accept(v)
	}
}

// TestVisitorPattern 单元测试
func TestVisitorPattern(t *testing.T) {
	// 1. Arrange: 准备对象结构 (Elements)
	collection := &ShapeCollection{}
	collection.Add(&Circle{Radius: 10})             // Area ≈ 314.159, Perimeter ≈ 62.83
	collection.Add(&Rectangle{Width: 2, Height: 3}) // Area = 6, Perimeter = 10

	// 2. Act & Assert: 使用“面积计算”访问者
	t.Run("CalculateArea", func(t *testing.T) {
		areaVisitor := &AreaCalculator{}

		// 触发访问
		collection.Accept(areaVisitor)

		expectedArea := (math.Pi * 10 * 10) + (2 * 3)
		if math.Abs(areaVisitor.TotalArea-expectedArea) > 0.001 {
			t.Errorf("面积计算错误。期望: %.2f, 实际: %.2f", expectedArea, areaVisitor.TotalArea)
		}
	})

	// 3. Act & Assert: 使用“周长计算”访问者
	t.Run("CalculatePerimeter", func(t *testing.T) {
		perimeterVisitor := &PerimeterCalculator{}

		// 触发访问
		collection.Accept(perimeterVisitor)

		expectedPerimeter := (2 * math.Pi * 10) + (2 * (2 + 3))
		if math.Abs(perimeterVisitor.TotalPerimeter-expectedPerimeter) > 0.001 {
			t.Errorf("周长计算错误。期望: %.2f, 实际: %.2f", expectedPerimeter, perimeterVisitor.TotalPerimeter)
		}
	})
}
