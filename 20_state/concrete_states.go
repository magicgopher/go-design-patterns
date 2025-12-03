package state

import (
	"errors"
	"fmt"
)

// NoItemState 无货状态
type NoItemState struct {
	vendingMachine *VendingMachine
}

// RequestItem 选中商品
func (i *NoItemState) RequestItem() error {
	return errors.New("缺货中，请勿选择")
}

// AddItem 添加商品
func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	// 状态流转：无货 -> 有货
	i.vendingMachine.SetState(i.vendingMachine.hasItemState)
	return nil
}

// InsertMoney 投入货币
func (i *NoItemState) InsertMoney(money int) error {
	return errors.New("缺货中，无法投币")
}

// DispenseItem 发放商品
func (i *NoItemState) DispenseItem() error {
	return errors.New("缺货中，无法出货")
}

// HasItemState 有货/空闲状态
type HasItemState struct {
	vendingMachine *VendingMachine
}

// RequestItem 选中商品
func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItemState)
		return errors.New("商品已售罄")
	}
	fmt.Println("商品已选中")
	// 状态流转：有货 -> 请求中
	i.vendingMachine.SetState(i.vendingMachine.itemRequestedState)
	return nil
}

// AddItem 添加商品
func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("增加商品 %d 个\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

// InsertMoney 投入货币
func (i *HasItemState) InsertMoney(money int) error {
	return errors.New("请先选择商品")
}

// DispenseItem 发放商品
func (i *HasItemState) DispenseItem() error {
	return errors.New("请先选择商品并支付")
}

// ItemRequestedState 商品请求/等待支付状态
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

// RequestItem 选中商品
func (i *ItemRequestedState) RequestItem() error {
	return errors.New("已经选择商品，请支付")
}

// AddItem 添加商品
func (i *ItemRequestedState) AddItem(count int) error {
	return errors.New("正在交易中，无法补货")
}

// InsertMoney 投入货币
func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("金额不足，需要 %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("投币成功")
	// 状态流转：请求中 -> 已付款
	i.vendingMachine.SetState(i.vendingMachine.hasMoneyState)
	return nil
}

// DispenseItem 发放商品
func (i *ItemRequestedState) DispenseItem() error {
	return errors.New("请先支付")
}

// HasMoneyState 已付款/出货状态
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

// RequestItem 选中商品
func (i *HasMoneyState) RequestItem() error {
	return errors.New("正在出货中")
}

// AddItem 添加商品
func (i *HasMoneyState) AddItem(count int) error {
	return errors.New("正在出货中，无法补货")
}

// InsertMoney 投入货币
func (i *HasMoneyState) InsertMoney(money int) error {
	return errors.New("已经支付，请等待出货")
}

// DispenseItem 发放商品
func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("正在发放商品...")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1

	// 状态流转：根据剩余库存决定回到 有货 还是 无货
	if i.vendingMachine.itemCount > 0 {
		i.vendingMachine.SetState(i.vendingMachine.hasItemState)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.noItemState)
	}
	return nil
}
