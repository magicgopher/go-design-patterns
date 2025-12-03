package strategy

// PaymentStrategy 定义支付策略接口
// 所有具体的支付方式都必须实现该接口
type PaymentStrategy interface {
	// Pay 执行支付逻辑
	Pay(amount float64) error
}
