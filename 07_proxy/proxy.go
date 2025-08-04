package proxy

import "fmt"

// 代理模式 - 代理

// Proxy 代理类，实现经销商的销售逻辑
type Proxy struct {
	seller *Seller
}

// NewProxy 创建新的代理实例
func NewProxy() *Proxy {
	return &Proxy{
		seller: &Seller{},
	}
}

// SaleComputer 代理销售电脑，增加20%利润
func (p *Proxy) SaleComputer(money float64) {
	profit := money * 0.2
	total := money + profit
	fmt.Printf("经销商代理销售电脑，原始价格: %.2f，增加20%%利润: %.2f，总价: %.2f\n", money, profit, total)
	p.seller.SaleComputer(money)
}
