package interpreter

import "testing"

// 客户端
// 单元测试

// TestInterpreterPattern 解释器模式单元测试
func TestInterpreterPattern(t *testing.T) {
	// 1. Arrange: 构建抽象语法树 (AST)
	// 目标表达式: a + b - c

	// 定义变量节点
	varA := NewVariable("a")
	varB := NewVariable("b")
	varC := NewVariable("c")

	// 构建运算树：(a + b) - c
	// 首先构建 (a + b)
	addExpr := NewAdd(varA, varB)
	// 然后构建 结果 - c
	finalExpr := NewSubtract(addExpr, varC)

	// 2. Arrange: 准备上下文环境 (Context)
	context := map[string]int{
		"a": 10,
		"b": 5,
		"c": 2,
	}

	// 3. Act: 执行解释
	result := finalExpr.Interpret(context)

	// 4. Assert: 验证结果
	// 10 + 5 - 2 = 13
	expected := 13
	if result != expected {
		t.Errorf("计算错误。期望: %d, 实际: %d", expected, result)
	}

	// 测试不同的上下文
	t.Run("DifferentContext", func(t *testing.T) {
		ctx2 := map[string]int{"a": 100, "b": 20, "c": 50}
		res2 := finalExpr.Interpret(ctx2) // 100 + 20 - 50 = 70
		if res2 != 70 {
			t.Errorf("Context2 计算错误。期望: 70, 实际: %d", res2)
		}
	})
}
