package adapter

import (
	"fmt"
	"strings"
)

// 适配器模式 - 适配器

// JSONToXMLAdapter 适配器，将 XML 支付接口适配为 JSON 支付系统
type JSONToXMLAdapter struct {
	legacySystem *LegacyPaymentSystem
}

// NewJSONToXMLAdapter 创建适配器实例
func NewJSONToXMLAdapter(legacySystem *LegacyPaymentSystem) *JSONToXMLAdapter {
	return &JSONToXMLAdapter{legacySystem: legacySystem}
}

// ProcessPayment 实现 Payment 接口，将 XML 数据转换为 JSON 并调用旧版系统
func (a *JSONToXMLAdapter) ProcessPayment(xmlData string) (string, error) {
	// 模拟 XML 到 JSON 的转换（简化为字符串替换）
	jsonData := strings.ReplaceAll(xmlData, "xml", "json")

	// 调用旧版支付系统的 JSON 接口
	result, err := a.legacySystem.ProcessJSONPayment(jsonData)
	if err != nil {
		return "", fmt.Errorf("adapter failed to process payment: %v", err)
	}

	// 模拟将 JSON 结果转换回 XML 格式
	xmlResult := strings.ReplaceAll(result, "JSON", "XML")
	return xmlResult, nil
}
