package proxy

// 代理模式 - 抽象主题

// Sale 销售接口
// 定义了代理和真实主题的公共接口，客户端通过此接口与代理或真实主题交互
type Sale interface {
	// SaleComputer 销售电脑
	// money 价格
	SaleComputer(money float64)
}
