package template

import (
	"fmt"
	"strings"
)

// SmsOtp 短信验证码实现
type SmsOtp struct {
	// 用于测试验证数据的字段
	LastCode string
	LastMsg  string
}

// GenRandomCode 生成随机验证码
func (s *SmsOtp) GenRandomCode(len int) string {
	// 简单模拟生成数字
	code := strings.Repeat("1", len)
	fmt.Printf("[SMS] 生成验证码: %s\n", code)
	s.LastCode = code
	return code
}

// SaveToCache 将验证码保存到缓存
func (s *SmsOtp) SaveToCache(code string) {
	fmt.Printf("[SMS] 缓存验证码 %s 到 Redis\n", code)
}

// GetMessage 组装短信消息
func (s *SmsOtp) GetMessage(code string) string {
	return fmt.Sprintf("【云服务】您的验证码是%s，5分钟内有效。", code)
}

// SendNotification 发送短信
func (s *SmsOtp) SendNotification(msg string) error {
	fmt.Printf("[SMS] 发送短信: %s\n", msg)
	s.LastMsg = msg
	return nil
}

// EmailOtp 邮件验证码实现
type EmailOtp struct {
	SenderEmail string
	LastMsg     string
}

// GenRandomCode 生成随机验证码
func (e *EmailOtp) GenRandomCode(len int) string {
	// 简单模拟生成字母
	code := strings.Repeat("A", len)
	fmt.Printf("[Email] 生成验证码: %s\n", code)
	return code
}

// SaveToCache 将验证码保存到缓存
func (e *EmailOtp) SaveToCache(code string) {
	fmt.Printf("[Email] 缓存验证码 %s 到 Memcached\n", code)
}

// GetMessage 组装邮件消息
func (e *EmailOtp) GetMessage(code string) string {
	return fmt.Sprintf("Subject: Login Code\nContent: Your code is %s", code)
}

// SendNotification 发送邮件
func (e *EmailOtp) SendNotification(msg string) error {
	fmt.Printf("[Email] 发送邮件 from %s: \n%s\n", e.SenderEmail, msg)
	e.LastMsg = msg
	return nil
}
