package visitor

// Shape 元素接口
// 定义 Accept 方法，允许访问者访问自己
type Shape interface {
	Accept(v Visitor)
}

// Visitor 访问者接口
// Go 不支持方法重载，因此必须为每种具体的 Element 定义明确的 Visit 方法
type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}
