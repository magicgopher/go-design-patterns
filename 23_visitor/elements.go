package visitor

// Circle 具体元素：圆形
type Circle struct {
	Radius float64
}

// Accept 实现 Shape 接口
// 核心逻辑：将自身（*Circle）传递给访问者的 VisitCircle 方法
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

// Rectangle 具体元素：矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// Accept 实现 Shape 接口
func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}
