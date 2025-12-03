package strategy

import "fmt"

// CreditCardStrategy 信用卡支付
type CreditCardStrategy struct {
	cardNumber string
	cvv        string
}

// NewCreditCardStrategy 创建信用卡支付策略
func NewCreditCardStrategy(cardNumber, cvv string) *CreditCardStrategy {
	return &CreditCardStrategy{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

// Pay 执行支付逻辑
func (c *CreditCardStrategy) Pay(amount float64) error {
	// 模拟信用卡支付逻辑
	fmt.Printf("已使用信用卡 %s 支付 %02.f 元\n", c.cardNumber, amount)
	return nil
}

// PayPalStrategy PayPal支付
type PayPalStrategy struct {
	email    string
	password string
}

// NewPayPalStrategy 创建PayPal支付策略
func NewPayPalStrategy(email, password string) *PayPalStrategy {
	return &PayPalStrategy{
		email:    email,
		password: password,
	}
}

// Pay 执行支付逻辑
func (p *PayPalStrategy) Pay(amount float64) error {
	// 模拟 PayPal 支付逻辑
	fmt.Printf("已使用 PayPal 账号 %s 支付 %02.f 元\n", p.email, amount)
	return nil
}

// WeChatPayStrategy 微信支付
type WeChatPayStrategy struct {
	openId string
}

// NewWeChatPayStrategy 创建微信支付策略
func NewWeChatPayStrategy(openId string) *WeChatPayStrategy {
	return &WeChatPayStrategy{openId: openId}
}

// Pay 执行支付逻辑
func (w *WeChatPayStrategy) Pay(amount float64) error {
	fmt.Printf("已使用微信支付 (OpenID: %s) 支付 %02.f 元\n", w.openId, amount)
	return nil
}
