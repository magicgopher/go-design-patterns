package state

import "fmt"

// VendingMachine 环境类：自动售货机
// 它维护当前状态，并将操作委托给当前状态处理
type VendingMachine struct {
	hasItemState       State
	itemRequestedState State
	hasMoneyState      State
	noItemState        State

	currentState State // 当前状态

	itemCount int // 商品数量
	itemPrice int // 商品价格
}

// NewVendingMachine 创建一个新的售货机实例
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	// 初始化所有具体状态，并将当前实例注入到状态中，以便状态能修改 Context
	v.hasItemState = &HasItemState{vendingMachine: v}
	v.itemRequestedState = &ItemRequestedState{vendingMachine: v}
	v.hasMoneyState = &HasMoneyState{vendingMachine: v}
	v.noItemState = &NoItemState{vendingMachine: v}

	// 设置初始状态
	if itemCount > 0 {
		v.currentState = v.hasItemState
	} else {
		v.currentState = v.noItemState
	}

	return v
}

// SetState 改变当前状态
func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

// 委托方法：将具体请求委托给当前状态处理

// RequestItem 选中商品
func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

// AddItem 添加商品
func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

// InsertMoney 投入货币
func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

// DispenseItem 发放商品
func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

// IncrementItemCount 增加商品库存（辅助方法）
func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("增加商品 %d 个\n", count)
	v.itemCount += count
}
