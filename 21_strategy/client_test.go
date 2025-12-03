package strategy

import "testing"

// TestStrategyPattern 模拟客户端选择不同的支付方式进行支付
func TestStrategyPattern(t *testing.T) {
	// 1. 模拟订单金额
	var amount float64 = 100

	// 2. 场景一：用户选择信用卡支付
	t.Run("PayWithCreditCard", func(t *testing.T) {
		// 创建具体策略
		ccStrategy := NewCreditCardStrategy("1234-5678-9012-3456", "123")

		// 创建 Context 并注入策略
		ctx := NewPaymentContext(ccStrategy)

		// 执行支付
		err := ctx.ExecutePayment(amount)
		if err != nil {
			t.Errorf("信用卡支付失败: %v", err)
		}
	})

	// 3. 场景二：用户切换为 PayPal 支付
	t.Run("SwitchToPayPal", func(t *testing.T) {
		// 创建初始 Context (假设默认是信用卡)
		ctx := NewPaymentContext(NewCreditCardStrategy("0000", "000"))

		// 用户决定改用 PayPal
		payPalStrategy := NewPayPalStrategy("user@example.com", "secret")

		// 动态切换策略
		ctx.SetStrategy(payPalStrategy)

		// 执行支付
		err := ctx.ExecutePayment(200)
		if err != nil {
			t.Errorf("PayPal 支付失败: %v", err)
		}
	})

	// 4. 场景三：未设置策略
	t.Run("NoStrategySet", func(t *testing.T) {
		// 创建一个空的 Context (Go 允许这样，虽然构造函数规避了，但直接初始化可能会发生)
		ctx := &PaymentContext{}

		err := ctx.ExecutePayment(100)
		if err == nil {
			t.Error("预期报错'未设置支付策略'，但未报错")
		}
	})
}
