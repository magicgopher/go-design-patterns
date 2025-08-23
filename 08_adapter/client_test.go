package adapter

import (
	"strings"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestModernPaymentSystem 测试现代支付系统直接处理 XML 数据的功能
func TestModernPaymentSystem(t *testing.T) {
	// 创建现代支付系统实例
	modernSystem := &ModernPaymentSystem{}

	// 定义测试用例的输入 XML 数据
	input := `<payment>xml_data</payment>`

	// 调用现代支付系统的 ProcessXMLPayment 方法
	result, err := modernSystem.ProcessXMLPayment(input)
	if err != nil {
		t.Errorf("现代支付系统处理失败: %v", err)
	}

	// 验证返回结果是否符合预期
	expected := "Processed XML payment: " + input
	if result != expected {
		t.Errorf("预期结果为 %q, 实际结果为 %q", expected, result)
	}
}

// TestJSONToXMLAdapter 测试适配器将 XML 数据适配到旧版 JSON 支付系统的功能
func TestJSONToXMLAdapter(t *testing.T) {
	// 创建旧版支付系统实例
	legacySystem := &LegacyPaymentSystem{}

	// 创建适配器实例，将旧版 JSON 支付系统适配为 XML 接口
	adapter := NewJSONToXMLAdapter(legacySystem)

	// 定义测试用例的输入 XML 数据
	input := `<payment>xml_data</payment>`

	// 通过适配器调用 ProcessPayment 方法
	result, err := adapter.ProcessPayment(input)
	if err != nil {
		t.Errorf("适配器处理失败: %v", err)
	}

	// 验证返回结果是否符合预期
	// 预期结果是旧版系统处理后的 JSON 数据被转换回 XML 格式
	expected := "Processed XML payment: " + strings.ReplaceAll(input, "xml", "json")
	if result != expected {
		t.Errorf("预期结果为 %q, 实际结果为 %q", expected, result)
	}
}

// TestJSONToXMLAdapterWithInvalidInput 测试适配器处理无效输入的情况
func TestJSONToXMLAdapterWithInvalidInput(t *testing.T) {
	// 创建旧版支付系统实例
	legacySystem := &LegacyPaymentSystem{}

	// 创建适配器实例
	adapter := NewJSONToXMLAdapter(legacySystem)

	// 定义无效的输入 XML 数据（转换后不包含 "json"）
	input := `<payment>invalid_data</payment>`

	// 通过适配器调用 ProcessPayment 方法，预期返回错误
	_, err := adapter.ProcessPayment(input)
	if err == nil {
		t.Error("预期适配器返回错误，但实际未返回错误")
	}

	// 验证错误信息是否包含预期的错误描述
	if !strings.Contains(err.Error(), "invalid JSON data") {
		t.Errorf("预期错误信息包含 'invalid JSON data', 实际错误为 %q", err)
	}
}
