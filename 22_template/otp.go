package template

import "fmt"

// IOtp 定义 OTP 流程中需要实现的具体步骤接口
type IOtp interface {
	GenRandomCode(len int) string
	SaveToCache(code string)
	GetMessage(code string) string
	SendNotification(msg string) error
}

// OtpWorkFlow 模板结构体
// 它持有 IOtp 接口，并定义了固定的执行流程
type OtpWorkFlow struct {
	iOtp IOtp
}

// NewOtpWorkFlow 构造函数
func NewOtpWorkFlow(i IOtp) *OtpWorkFlow {
	return &OtpWorkFlow{iOtp: i}
}

// GenAndSendOTP 模板方法 (Template Method)
// 这里定义了算法的骨架：生成 -> 缓存 -> 组装消息 -> 发送
// 这个流程是固定的，但具体的实现委托给了 iOtp
func (o *OtpWorkFlow) GenAndSendOTP(otpLength int) error {
	// 步骤 1: 生成
	code := o.iOtp.GenRandomCode(otpLength)

	// 步骤 2: 缓存
	o.iOtp.SaveToCache(code)

	// 步骤 3: 组装
	msg := o.iOtp.GetMessage(code)

	// 步骤 4: 发送
	err := o.iOtp.SendNotification(msg)
	if err != nil {
		return err
	}

	fmt.Println("流程结束：发送成功")
	return nil
}
