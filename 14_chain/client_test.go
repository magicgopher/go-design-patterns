package chain

import (
	"fmt"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestManagerApproval 测试 Manager 是否能正确处理金额小于 1000 的采购请求
func TestManagerApproval(t *testing.T) {
	manager := &Manager{}
	request := &PurchaseRequest{Amount: 500, Purpose: "办公用品"}
	expected := fmt.Sprintf("经理批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
	result := manager.ProcessRequest(request)
	if result != expected {
		t.Errorf("预期结果: %s, 实际结果: %s", expected, result)
	}
	t.Logf("Manager 成功处理金额 %f 的采购请求，用途: %s", request.Amount, request.Purpose)
}

// TestDirectorApproval 测试 Director 是否能正确处理金额在 1000 到 10000 之间的采购请求
func TestDirectorApproval(t *testing.T) {
	manager := &Manager{}
	director := &Director{}
	manager.SetNext(director)
	request := &PurchaseRequest{Amount: 5000, Purpose: "设备采购"}
	expected := fmt.Sprintf("总监批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
	result := manager.ProcessRequest(request)
	if result != expected {
		t.Errorf("预期结果: %s, 实际结果: %s", expected, result)
	}
	t.Logf("Director 成功处理金额 %f 的采购请求，用途: %s", request.Amount, request.Purpose)
}

// TestCEOApproval 测试 CEO 是否能正确处理金额大于 10000 的采购请求
func TestCEOApproval(t *testing.T) {
	manager := &Manager{}
	director := &Director{}
	ceo := &CEO{}
	manager.SetNext(director)
	director.SetNext(ceo)
	request := &PurchaseRequest{Amount: 15000, Purpose: "大型项目投资"}
	expected := fmt.Sprintf("CEO 批准了采购 '%s'，金额：%f", request.Purpose, request.Amount)
	result := manager.ProcessRequest(request)
	if result != expected {
		t.Errorf("预期结果: %s, 实际结果: %s", expected, result)
	}
	t.Logf("CEO 成功处理金额 %f 的采购请求，用途: %s", request.Amount, request.Purpose)
}

// TestNoHandler 测试当链中没有合适的处理者时是否返回正确提示
func TestNoHandler(t *testing.T) {
	manager := &Manager{}
	request := &PurchaseRequest{Amount: 15000, Purpose: "超额采购"}
	expected := "无人能处理此采购请求"
	result := manager.ProcessRequest(request)
	if result != expected {
		t.Errorf("预期结果: %s, 实际结果: %s", expected, result)
	}
	t.Logf("无合适处理者，金额 %f 的采购请求被拒绝，用途: %s", request.Amount, request.Purpose)
}

// TestChainOfResponsibility 测试请求是否能正确沿责任链传递
func TestChainOfResponsibility(t *testing.T) {
	manager := &Manager{}
	director := &Director{}
	ceo := &CEO{}
	manager.SetNext(director)
	director.SetNext(ceo)

	tests := []struct {
		amount   float64
		purpose  string
		expected string
	}{
		{500, "办公用品", fmt.Sprintf("经理批准了采购 '办公用品'，金额：%f", 500.0)},
		{5000, "设备采购", fmt.Sprintf("总监批准了采购 '设备采购'，金额：%f", 5000.0)},
		{15000, "大型项目投资", fmt.Sprintf("CEO 批准了采购 '大型项目投资'，金额：%f", 15000.0)},
	}

	for _, test := range tests {
		request := &PurchaseRequest{Amount: test.amount, Purpose: test.purpose}
		result := manager.ProcessRequest(request)
		if result != test.expected {
			t.Errorf("金额: %f, 用途: %s, 预期结果: %s, 实际结果: %s", test.amount, test.purpose, test.expected, result)
		}
		t.Logf("责任链成功处理金额 %f 的采购请求，用途: %s", test.amount, test.purpose)
	}
}
