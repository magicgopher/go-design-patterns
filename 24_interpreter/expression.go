package interpreter

// Expression 表示表达式接口
// 定义了解释操作，Context 在这里体现为一个简单的 map，用于存储变量值
type Expression interface {
	Interpret(context map[string]int) int
}
