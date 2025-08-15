package adapter

import (
	"fmt"
	"strings"
)

// 适配器模式 - 适配器

// XMLToJSONAdapter 适配器，将XML支付系统适配为JSON支付接口
type XMLToJSONAdapter struct {
	legacySystem *LegacyPaymentSystem
}

// NewXMLToJSONAdapter 创建适配器实例
func NewXMLToJSONAdapter(legacySystem *LegacyPaymentSystem) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{legacySystem: legacySystem}
}

// ProcessPayment 实现Payment接口，将JSON数据转换为XML并调用旧版系统
func (a *XMLToJSONAdapter) ProcessPayment(jsonData string) (string, error) {
	// 模拟JSON到XML的转换（简化为字符串替换）
	xmlData := strings.ReplaceAll(jsonData, "json", "xml")

	// 调用旧版支付系统的XML接口
	result, err := a.legacySystem.ProcessXMLPayment(xmlData)
	if err != nil {
		return "", fmt.Errorf("adapter failed to process payment: %v", err)
	}

	// 模拟将XML结果转换回JSON格式
	jsonResult := strings.ReplaceAll(result, "XML", "JSON")
	return jsonResult, nil
}
