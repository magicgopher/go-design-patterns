package adapter

import "fmt"

// 适配器模式 - 新版支付系统（JSON）

// ModernPaymentSystem 模拟新版支付系统，处理 XML 数据
type ModernPaymentSystem struct{}

// ProcessXMLPayment 处理 XML 格式的支付请求
func (m *ModernPaymentSystem) ProcessXMLPayment(xmlData string) (string, error) {
	return fmt.Sprintf("Processed XML payment: %s", xmlData), nil
}
