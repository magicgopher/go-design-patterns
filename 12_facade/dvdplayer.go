package facade

import "fmt"

// 外观模式

// 子系统角色

// DVDPlayer 是子系统组件，负责 DVD 播放
type DVDPlayer struct{}

func (d *DVDPlayer) TurnOn() {
	fmt.Println("DVD 播放器已开启")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("正在播放电影: %s\n", movie)
}

func (d *DVDPlayer) TurnOff() {
	fmt.Println("DVD 播放器已关闭")
}
