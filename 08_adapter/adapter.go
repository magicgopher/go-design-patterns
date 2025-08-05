package adapter

import (
	"fmt"
	"strings"
)

// 适配器模式 - 适配器

// JSONToXMLAdapter 适配器，将JSON支付系统适配为XML支付接口
type JSONToXMLAdapter struct {
	modernSystem *ModernPaymentSystem
}

// NewJSONToXMLAdapter 创建适配器实例
func NewJSONToXMLAdapter(modernSystem *ModernPaymentSystem) *JSONToXMLAdapter {
	return &JSONToXMLAdapter{modernSystem: modernSystem}
}

// ProcessPayment 实现Payment接口，将XML数据转换为JSON并调用新版系统
func (a *JSONToXMLAdapter) ProcessPayment(xmlData string) (string, error) {
	// 模拟XML到JSON的转换（简化为字符串替换）
	jsonData := strings.Replace(xmlData, "xml", "json", -1)

	// 调用新版支付系统的JSON接口
	result, err := a.modernSystem.ProcessJSONPayment(jsonData)
	if err != nil {
		return "", fmt.Errorf("adapter failed to process payment: %v", err)
	}

	// 模拟将JSON结果转换回XML格式
	xmlResult := strings.Replace(result, "JSON", "XML", -1)
	return xmlResult, nil
}
