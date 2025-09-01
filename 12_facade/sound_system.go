package facade

import "fmt"

// 外观模式

// 子系统角色

// SoundSystem 是子系统组件，负责音响系统
type SoundSystem struct{}

func (s *SoundSystem) TurnOn() {
	fmt.Println("音响系统已开启")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("音响音量设置为: %d\n", level)
}

func (s *SoundSystem) TurnOff() {
	fmt.Println("音响系统已关闭")
}
