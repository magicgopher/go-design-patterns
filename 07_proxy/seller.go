package proxy

import "fmt"

// 代理模式 - 真实主题

// Seller 经销商
// 真实主题角色，实现了 Sale 接口，负责实际的销售电脑逻辑
type Seller struct {
}

// SaleComputer 实现销售电脑的具体逻辑
// money 价格，表示客户支付的金额
func (s Seller) SaleComputer(money float64) {
	fmt.Printf("销售电脑, 并拿到钱:%.2f", money)
}
