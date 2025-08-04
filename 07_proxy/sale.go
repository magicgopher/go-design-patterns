package proxy

// 代理模式 - 抽象主题

// Sale 销售接口
type Sale interface {
	// SaleComputer 销售电脑
	// money 价格
	SaleComputer(money float64)
}
