package strategy

import "errors"

// PaymentContext 支付环境类
// 它持有当前的支付策略，并提供执行支付的方法
type PaymentContext struct {
	strategy PaymentStrategy
}

// NewPaymentContext 创建一个新的支付环境，需指定初始策略
func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		strategy: strategy,
	}
}

// SetStrategy 允许在运行时动态切换支付策略
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// ExecutePayment 委托当前策略执行支付
func (p *PaymentContext) ExecutePayment(amount float64) error {
	if p.strategy == nil {
		return errors.New("未设置支付策略")
	}
	return p.strategy.Pay(amount)
}
