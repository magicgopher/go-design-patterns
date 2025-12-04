package template

import (
	"strings"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestTemplateMethod 单元测试
// 验证通过模板方法调用不同实现时的行为是否正确
func TestTemplateMethod(t *testing.T) {

	// 场景一：测试 SMS 流程
	t.Run("SMS_Flow", func(t *testing.T) {
		// Arrange
		smsOtp := &SmsOtp{}
		workflow := NewOtpWorkFlow(smsOtp)

		// Act
		err := workflow.GenAndSendOTP(4)

		// Assert
		if err != nil {
			t.Fatalf("SMS 流程执行失败: %v", err)
		}

		// 验证具体步骤是否被正确调用（检查副作用）
		expectedCode := "1111"
		if smsOtp.LastCode != expectedCode {
			t.Errorf("SMS 验证码生成错误。期望: %s, 实际: %s", expectedCode, smsOtp.LastCode)
		}

		if !strings.Contains(smsOtp.LastMsg, "【云服务】") {
			t.Errorf("SMS 消息格式错误。实际内容: %s", smsOtp.LastMsg)
		}
	})

	// 场景二：测试 Email 流程
	t.Run("Email_Flow", func(t *testing.T) {
		// Arrange
		emailOtp := &EmailOtp{SenderEmail: "admin@test.com"}
		workflow := NewOtpWorkFlow(emailOtp)

		// Act
		err := workflow.GenAndSendOTP(6)

		// Assert
		if err != nil {
			t.Fatalf("Email 流程执行失败: %v", err)
		}

		// 验证 Email 特有的逻辑
		expectedContent := "Your code is AAAAAA"
		if !strings.Contains(emailOtp.LastMsg, expectedContent) {
			t.Errorf("Email 内容错误。期望包含: %s, 实际: %s", expectedContent, emailOtp.LastMsg)
		}
	})
}
