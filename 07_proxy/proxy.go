package proxy

import "fmt"

// 代理模式 - 代理

// Proxy 代理
// 代理角色，实现了 Sale 接口，控制对 Seller 的访问并添加额外逻辑
type Proxy struct {
	seller *Seller // 持有对真实主题（Seller）的引用
}

// NewProxy 创建新的代理实例
// 返回一个初始化了 Seller 的 Proxy 实例
func NewProxy() *Proxy {
	return &Proxy{
		seller: &Seller{}, // 初始化真实主题
	}
}

// SaleComputer 代理销售电脑，增加20%利润
// money 原始价格，代理会在此基础上计算利润并调用真实主题的逻辑
func (p *Proxy) SaleComputer(money float64) {
	// 计算20%的利润
	profit := money * 0.2
	// 计算总价（原始价格 + 利润）
	total := money + profit
	// 打印代理的销售信息，包括原始价格、利润和总价
	fmt.Printf("经销商代理销售电脑，原始价格: %.2f，增加20%%利润: %.2f，总价: %.2f\n", money, profit, total)
	// 调用真实主题的销售逻辑
	p.seller.SaleComputer(money)
}
