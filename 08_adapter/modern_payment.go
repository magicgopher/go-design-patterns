package adapter

import "fmt"

// 适配器模式 - 新版支付系统（JSON）被适配者

// ModernPaymentSystem 模拟新版支付系统，处理JSON数据
type ModernPaymentSystem struct{}

// ProcessJSONPayment 处理JSON格式的支付请求
func (m *ModernPaymentSystem) ProcessJSONPayment(jsonData string) (string, error) {
	return fmt.Sprintf("Processed JSON payment: %s", jsonData), nil
}
