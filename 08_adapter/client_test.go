package adapter

import "testing"

// 单元测试
// 模拟客户端调用

// TestJSONToXMLAdapter 测试适配器是否能正确将XML请求适配为JSON并返回预期的XML结果
func TestJSONToXMLAdapter(t *testing.T) {
	// 创建新版支付系统实例
	modernSystem := &ModernPaymentSystem{}

	// 创建适配器实例，将新版支付系统传入
	adapter := NewJSONToXMLAdapter(modernSystem)

	// 定义测试用例，包含输入的XML数据、期望的输出结果和是否期望错误
	tests := []struct {
		inputXML    string // 输入的XML格式数据
		expected    string // 期望的输出结果（XML格式）
		expectError bool   // 是否期望发生错误
	}{
		{
			inputXML:    "<payment>xml data</payment>",                         // 测试用例1：输入XML数据
			expected:    "Processed XML payment: <payment>json data</payment>", // 期望的XML输出
			expectError: false,                                                 // 不期望发生错误
		},
		{
			inputXML:    "<payment>another xml data</payment>",                         // 测试用例2：输入另一组XML数据
			expected:    "Processed XML payment: <payment>another json data</payment>", // 期望的XML输出
			expectError: false,                                                         // 不期望发生错误
		},
	}

	// 遍历测试用例，执行测试
	for _, test := range tests {
		t.Run(test.inputXML, func(t *testing.T) {
			// 调用适配器的ProcessPayment方法，处理XML输入
			result, err := adapter.ProcessPayment(test.inputXML)

			// 检查错误是否符合预期
			if (err != nil) != test.expectError {
				t.Errorf("Expected error: %v, got: %v", test.expectError, err)
			}

			// 检查输出结果是否符合预期
			if result != test.expected {
				t.Errorf("Expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
