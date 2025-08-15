package adapter

import "testing"

// 单元测试
// 模拟客户端调用

// TestXMLToJSONAdapter 测试适配器是否能正确将JSON请求适配为XML并返回预期的JSON结果
func TestXMLToJSONAdapter(t *testing.T) {
	// 创建旧版支付系统实例
	legacySystem := &LegacyPaymentSystem{}

	// 创建适配器实例，将旧版支付系统传入
	adapter := NewXMLToJSONAdapter(legacySystem)

	// 定义测试用例，包含输入的JSON数据、期望的输出结果和是否期望错误
	tests := []struct {
		inputJSON   string // 输入的JSON格式数据
		expected    string // 期望的输出结果（JSON格式）
		expectError bool   // 是否期望发生错误
	}{
		{
			inputJSON:   `{"payment":"json data"}`,                        // 测试用例1：输入JSON数据
			expected:    `Processed JSON payment: {"payment":"xml data"}`, // 期望的JSON输出
			expectError: false,                                            // 不期望发生错误
		},
		{
			inputJSON:   `{"payment":"another json data"}`,                        // 测试用例2：输入另一组JSON数据
			expected:    `Processed JSON payment: {"payment":"another xml data"}`, // 期望的JSON输出
			expectError: false,                                                    // 不期望发生错误
		},
	}

	// 遍历测试用例，执行测试
	for _, test := range tests {
		t.Run(test.inputJSON, func(t *testing.T) {
			// 调用适配器的ProcessPayment方法，处理JSON输入
			result, err := adapter.ProcessPayment(test.inputJSON)

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
