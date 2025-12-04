package visitor

import "math"

// 具体的访问者实现

// 面积计算访问者

// AreaCalculator 计算面积
type AreaCalculator struct {
	TotalArea float64 // 用于累加计算结果
}

// VisitCircle 计算圆形面积
func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := math.Pi * c.Radius * c.Radius
	a.TotalArea += area
}

// VisitRectangle 计算矩形面积
func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	a.TotalArea += area
}

// 周长计算访问者

// PerimeterCalculator 计算周长
type PerimeterCalculator struct {
	TotalPerimeter float64
}

// VisitCircle 计算圆形周长
func (p *PerimeterCalculator) VisitCircle(c *Circle) {
	perimeter := 2 * math.Pi * c.Radius
	p.TotalPerimeter += perimeter
}

// VisitRectangle 计算矩形周长
func (p *PerimeterCalculator) VisitRectangle(r *Rectangle) {
	perimeter := 2 * (r.Width + r.Height)
	p.TotalPerimeter += perimeter
}
