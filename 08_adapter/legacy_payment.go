package adapter

import (
	"fmt"
	"strings"
)

// 适配器模式 - 旧版支付系统（XML）被适配者

// LegacyPaymentSystem 模拟旧版支付系统，处理 JSON 数据
type LegacyPaymentSystem struct{}

// ProcessJSONPayment 处理 JSON 格式的支付请求
func (l *LegacyPaymentSystem) ProcessJSONPayment(jsonData string) (string, error) {
	// 模拟简单的输入验证
	if !strings.Contains(jsonData, "json") {
		return "", fmt.Errorf("invalid JSON data: %s", jsonData)
	}
	return fmt.Sprintf("Processed JSON payment: %s", jsonData), nil
}
