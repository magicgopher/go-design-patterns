package command

import "testing"

// 单元测试
// 模拟客户端调用

// TestCommand 测试命令模式的功能
func TestCommand(t *testing.T) {
	// 创建接收者（客厅的灯）
	livingRoomLight := &Light{Name: "客厅"}

	// 创建具体命令并绑定接收者
	turnOn := NewTurnOnCommand(livingRoomLight)   // 创建开灯命令
	turnOff := NewTurnOffCommand(livingRoomLight) // 创建关灯命令

	// 创建调用者（开关）
	lightSwitch := &Switch{}

	t.Run("Turn On Light", func(t *testing.T) {
		// 配置调用者和开灯命令
		lightSwitch.SetCommand(turnOn)
		// 按下开关，执行开灯命令
		lightSwitch.Press()
	})

	t.Run("Turn Off Light", func(t *testing.T) {
		// 重新配置调用者和关灯命令
		lightSwitch.SetCommand(turnOff)
		// 按下开关，执行关灯命令
		lightSwitch.Press()
	})
}
