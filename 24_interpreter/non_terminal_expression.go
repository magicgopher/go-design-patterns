package interpreter

// Add 非终结符表达式：加法
// 它包含左右两个子表达式
type Add struct {
	left  Expression
	right Expression
}

// NewAdd 构造函数
func NewAdd(left, right Expression) *Add {
	return &Add{left: left, right: right}
}

// Interpret 递归计算左右表达式并相加
func (a *Add) Interpret(context map[string]int) int {
	return a.left.Interpret(context) + a.right.Interpret(context)
}

// Subtract 非终结符表达式：减法
type Subtract struct {
	left  Expression
	right Expression
}

// NewSubtract 构造函数
func NewSubtract(left, right Expression) *Subtract {
	return &Subtract{left: left, right: right}
}

// Interpret 递归计算左右表达式并相减
func (s *Subtract) Interpret(context map[string]int) int {
	return s.left.Interpret(context) - s.right.Interpret(context)
}
