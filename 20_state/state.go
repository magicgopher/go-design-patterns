package state

// 状态模式

// 定义状态接口

// State 定义状态接口
// 包含售货机所有可能触发的动作
type State interface {
	// AddItem 添加商品
	AddItem(count int) error

	// RequestItem 请求购买商品
	RequestItem() error

	// InsertMoney 投入货币
	InsertMoney(money int) error

	// DispenseItem 发放商品
	DispenseItem() error
}
