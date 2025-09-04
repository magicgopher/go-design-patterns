package command

// 命令模式

// Switch 是调用者，它持有一个命令并触发它
type Switch struct {
	command Command
}

// SetCommand 设置命令
func (s *Switch) SetCommand(cmd Command) {
	s.command = cmd
}

// Press 模拟按钮按下，执行命令
func (s *Switch) Press() {
	s.command.Execute()
}
