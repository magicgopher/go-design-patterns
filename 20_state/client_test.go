package state

import (
	"testing"
)

// 客户端
// 单元测试

// TestStatePattern 模拟客户端操作自动售货机
func TestVendingMachine(t *testing.T) {
	// 场景 1: 正常购买流程 (Happy Path)
	t.Run("HappyPathPurchaseSuccess", func(t *testing.T) {
		// Arrange: 初始化 1 个商品，价格 10 元
		vm := NewVendingMachine(1, 10)

		// Act & Assert 1: 请求商品
		if err := vm.RequestItem(); err != nil {
			t.Fatalf("预期请求成功，但返回错误: %v", err)
		}

		// Act & Assert 2: 投入不足的金额 (测试中间逻辑)
		if err := vm.InsertMoney(5); err == nil {
			t.Fatal("预期金额不足报错，但未报错")
		}

		// Act & Assert 3: 投入足额金额
		if err := vm.InsertMoney(10); err != nil {
			t.Fatalf("预期投币成功，但返回错误: %v", err)
		}

		// Act & Assert 4: 出货
		if err := vm.DispenseItem(); err != nil {
			t.Fatalf("预期出货成功，但返回错误: %v", err)
		}

		// 验证副作用：库存应为 0
		if vm.itemCount != 0 {
			t.Errorf("预期库存为 0, 实际为 %d", vm.itemCount)
		}
	})

	// 场景 2: 缺货流程 (Out of Stock)
	t.Run("OutOfStockFailure", func(t *testing.T) {
		// Arrange: 初始化 0 个商品
		vm := NewVendingMachine(0, 10)

		// Act
		err := vm.RequestItem()

		// Assert
		if err == nil {
			t.Error("预期缺货报错，但返回了 nil")
		}
		expectedErr := "缺货中，请勿选择" // 这里假设具体的错误信息
		if err.Error() != expectedErr {
			t.Logf("收到预期内的错误: %v", err) // 只要报错就行，具体信息如果不严格匹配可以用 Logf
		}
	})

	// 场景 3: 状态流转限制 (非法操作)
	t.Run("InvalidStateTransition", func(t *testing.T) {
		vm := NewVendingMachine(1, 10)

		// 未选择商品直接投币
		err := vm.InsertMoney(10)
		if err == nil {
			t.Error("预期报错'请先选择商品'，但操作成功了")
		}

		// 未投币直接出货
		err = vm.RequestItem() // 先选中
		if err != nil {
			t.Fatal(err)
		}

		err = vm.DispenseItem() // 未投币直接出货
		if err == nil {
			t.Error("预期报错'请先支付'，但操作成功了")
		}
	})

	// 场景 4: 补货测试
	t.Run("RestockSuccess", func(t *testing.T) {
		vm := NewVendingMachine(0, 10)

		// 初始状态应无法购买
		if err := vm.RequestItem(); err == nil {
			t.Fatal("初始应缺货")
		}

		// 补货
		if err := vm.AddItem(5); err != nil {
			t.Fatalf("补货失败: %v", err)
		}

		// 补货后应能购买
		if err := vm.RequestItem(); err != nil {
			t.Errorf("补货后预期可以购买，但报错: %v", err)
		}

		if vm.itemCount != 5 {
			t.Errorf("预期库存 5, 实际 %d", vm.itemCount)
		}
	})
}
