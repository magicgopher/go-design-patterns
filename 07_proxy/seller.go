package proxy

import "fmt"

// 代理模式 - 真实主题

type Seller struct {
}

func (s Seller) SaleComputer(money float64) {
	fmt.Printf("销售电脑, 并拿到钱:%.2f", money)
}
