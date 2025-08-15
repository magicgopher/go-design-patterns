package adapter

import "fmt"

// 适配器模式 - 旧版支付系统（XML）被适配者

// LegacyPaymentSystem 模拟旧版支付系统，处理XML数据
type LegacyPaymentSystem struct{}

// ProcessXMLPayment 处理XML格式的支付请求
func (l *LegacyPaymentSystem) ProcessXMLPayment(xmlData string) (string, error) {
	return fmt.Sprintf("Processed XML payment: %s", xmlData), nil
}
