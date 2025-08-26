package bridge

import (
	"fmt"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestSendCommonMessageBySms 模拟通过短信发送普通消息
func TestSendCommonMessageBySms(t *testing.T) {
	fmt.Println("--- 场景1: 通过短信发送普通消息 ---")
	// 1. 选择一个实现（短信发送器）
	sender := NewSmsSender()

	// 2. 选择一个抽象（普通消息），并将实现“桥接”进去
	commonMessage := NewCommonMessage(sender)

	// 3. 调用抽象的方法
	commonMessage.Send("你好，这是一个普通消息。")
}

// TestSendUrgencyMessageByEmail 模拟通过邮件发送加急消息
func TestSendUrgencyMessageByEmail(t *testing.T) {
	fmt.Println("\n--- 场景2: 通过邮件发送加急消息 ---")
	// 1. 选择一个实现（邮件发送器）
	sender := NewEmailSender()

	// 2. 选择一个抽象（加急消息），并将实现“桥接”进去
	urgencyMessage := NewUrgencyMessage(sender)

	// 3. 调用抽象的方法
	urgencyMessage.Send("系统出现严重故障，请立即处理！")
}

// TestSendUrgencyMessageBySms 模拟通过短信发送加急消息
func TestSendUrgencyMessageBySms(t *testing.T) {
	fmt.Println("\n--- 场景3: 通过短信发送加急消息 (更换实现) ---")
	// 1. 改变实现（短信发送器）
	sender := NewSmsSender()

	// 2. 抽象部分（加急消息）不变，仅需更换桥接的实现
	urgencyMessage := NewUrgencyMessage(sender)

	// 3. 调用抽象的方法
	urgencyMessage.Send("服务器即将宕机，十万火急！")
}
