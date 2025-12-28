package interpreter

// Variable 终结符表达式：变量
// 它代表语法树的叶子节点
type Variable struct {
	name string
}

// NewVariable 构造函数
func NewVariable(name string) *Variable {
	return &Variable{name: name}
}

// Interpret 从上下文中获取变量的值
func (v *Variable) Interpret(context map[string]int) int {
	if val, ok := context[v.name]; ok {
		return val
	}
	return 0 // 默认值或错误处理
}
