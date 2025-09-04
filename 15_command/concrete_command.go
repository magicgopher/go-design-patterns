package command

// 命令模式

// TurnOnCommand 是一个具体的命令，用于开灯
type TurnOnCommand struct {
	light *Light
}

// NewTurnOnCommand 创建 TurnOnCommand
func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

// Execute 执行命令
func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// TurnOffCommand 是一个具体的命令，用于关灯
type TurnOffCommand struct {
	light *Light
}

// NewTurnOffCommand 创建 TurnOffCommand
func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

// Execute 执行命令
func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}
