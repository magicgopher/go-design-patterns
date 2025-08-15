package adapter

// 适配器模式 - 目标接口

// Payment 定义目标接口，客户端期望使用此接口
type Payment interface {
	ProcessPayment(jsonData string) (string, error)
}
